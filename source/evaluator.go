package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EvaluateNodes(nodes []*Node, rt *RuntimeEnv) error {
	for _, n := range nodes {
		switch n.Kind {
		case NodeSay:
			fmt.Println(n.Expr)
			rt.mu.Lock()
			rt.Variables["last_say"] = Value{Kind: ValString, Str: n.Expr}
			rt.mu.Unlock()
		case NodeAssign:
			handleLet(n.Expr, rt)
		case NodeMath:
			v, err := EvalExpressionWithVars(n.Expr, rt.Variables)
			if err != nil {
				fmt.Println("Math error:", err)
			} else {
				printNumber(v)
			}
		case NodeMathBlock:
			for _, l := range n.Lines {
				if strings.TrimSpace(l) == "" {
					continue
				}
				v, err := EvalExpressionWithVars(l, rt.Variables)
				if err != nil {
					fmt.Println("Math error:", err)
				} else {
					printNumber(v)
				}
			}
		case NodeRun:
			if len(n.Args) >= 1 {
				_ = rt.RunScript(n.Args[0])
			}
		case NodeBuild:
			if len(n.Args) >= 1 {
				outPath := ""
				arch := ""
				if len(n.Args) >= 2 {
					outPath = n.Args[1]
				}
				if len(n.Args) >= 3 {
					arch = n.Args[2]
				}
				res, err := rt.BuildScript(n.Args[0], outPath, arch)
				if err != nil {
					fmt.Println("build error:", err)
				} else {
					fmt.Println(res)
				}
			}
		case NodeExec:
			if len(n.Args) >= 1 {
				_, _ = rt.ExecShell(strings.Join(n.Args, " "), strings.HasPrefix(n.Raw, "exec-safe"))
			}
		case NodeTimePrint:
			fmt.Println(PrintTime())
		case NodeTimeSet:
			if len(n.Args) >= 2 {
				SetGlobalTime(trimQuotes(n.Args[0]), trimQuotes(n.Args[1]))
			}
		case NodeIf:
			if evalConditionFromNode(n, rt) {
				_ = EvaluateNodes(n.Children, rt)
			}
		default:
			line := strings.TrimSpace(n.Raw)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			if strings.HasPrefix(line, "say ") {
				arg := strings.TrimSpace(strings.TrimPrefix(line, "say "))
				res := evalSayConcat(arg, rt)
				fmt.Println(res)
				rt.mu.Lock()
				rt.Variables["last_say"] = Value{Kind: ValString, Str: res}
				rt.mu.Unlock()
				continue
			}
			if strings.HasPrefix(line, "let ") {
				handleLet(strings.TrimSpace(strings.TrimPrefix(line, "let")), rt)
				continue
			}
			v, err := EvalExpressionWithVars(line, rt.Variables)
			if err == nil {
				printNumber(v)
				continue
			}
			fmt.Println("Unknown command:", n.Raw)
		}
	}
	return nil
}

func printNumber(v float64) {
	if v == float64(int64(v)) {
		fmt.Printf("%.0f\n", v)
	} else {
		fmt.Printf("%v\n", v)
	}
}

func evalSayConcat(arg string, rt *RuntimeEnv) string {
	parts := splitByPlus(arg)
	out := strings.Builder{}
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if (p[0] == '"' && p[len(p)-1] == '"') || (p[0] == '\'' && p[len(p)-1] == '\'') {
			out.WriteString(trimQuotes(p))
			continue
		}
		if val, ok := rt.Variables[p]; ok {
			if val.Kind == ValString {
				out.WriteString(val.Str)
				continue
			}
			if val.Kind == ValNumber {
				out.WriteString(strconv.FormatFloat(val.Num, 'f', -1, 64))
				continue
			}
		}
		v, err := EvalExpressionWithVars(p, rt.Variables)
		if err == nil {
			out.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
			continue
		}
		out.WriteString(p)
	}
	return out.String()
}

func splitByPlus(s string) []string {
	parts := []string{}
	cur := strings.Builder{}
	inQuotes := false
	var q rune
	for _, r := range s {
		if inQuotes {
			cur.WriteRune(r)
			if r == q {
				inQuotes = false
			}
			continue
		}
		if r == '"' || r == '\'' {
			inQuotes = true
			q = r
			cur.WriteRune(r)
			continue
		}
		if r == '+' {
			parts = append(parts, cur.String())
			cur.Reset()
			continue
		}
		cur.WriteRune(r)
	}
	if cur.Len() > 0 {
		parts = append(parts, cur.String())
	}
	return parts
}

func handleLet(expr string, rt *RuntimeEnv) {
	rest := strings.TrimSpace(expr)
	if idx := strings.Index(rest, "="); idx >= 0 {
		name := strings.TrimSpace(rest[:idx])
		right := strings.TrimSpace(rest[idx+1:])
		if right == "" {
			return
		}
		if (right[0] == '"' && right[len(right)-1] == '"') || (right[0] == '\'' && right[len(right)-1] == '\'') {
			rt.mu.Lock()
			rt.Variables[name] = Value{Kind: ValString, Str: trimQuotes(right)}
			rt.mu.Unlock()
			return
		}
		if v, err := EvalExpressionWithVars(right, rt.Variables); err == nil {
			rt.mu.Lock()
			rt.Variables[name] = Value{Kind: ValNumber, Num: v}
			rt.mu.Unlock()
			return
		}
		rt.mu.Lock()
		rt.Variables[name] = Value{Kind: ValString, Str: right}
		rt.mu.Unlock()
	}
}

func evalConditionFromNode(n *Node, rt *RuntimeEnv) bool {
	cond := strings.TrimSpace(n.Raw)
	cond = strings.TrimPrefix(cond, "if")
	cond = strings.TrimSpace(cond)
	return evalCondition(cond, rt)
}

func evalCondition(cond string, rt *RuntimeEnv) bool {
	ops := []string{">=", "<=", "==", "!=", ">", "<", "="}
	for _, op := range ops {
		if idx := strings.Index(cond, op); idx >= 0 {
			left := strings.TrimSpace(cond[:idx])
			right := strings.TrimSpace(cond[idx+len(op):])
			if (left != "" && (left[0] == '"' || left[0] == '\'')) || (right != "" && (right[0] == '"' || right[0] == '\'')) {
				l := trimQuotes(left)
				r := trimQuotes(right)
				switch op {
				case "==", "=", "!=":
					if op == "!=" {
						return l != r
					}
					return l == r
				}
				return false
			}
			ln, lerr := EvalExpressionWithVars(left, rt.Variables)
			rn, rerr := EvalExpressionWithVars(right, rt.Variables)
			if lerr != nil || rerr != nil {
				return false
			}
			switch op {
			case ">=":
				return ln >= rn
			case "<=":
				return ln <= rn
			case ">":
				return ln > rn
			case "<":
				return ln < rn
			case "==", "=":
				return ln == rn
			case "!=":
				return ln != rn
			}
		}
	}
	return false
}
