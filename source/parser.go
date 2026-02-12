package main

import (
	"os"
	"strings"
)

func ParseFileToNodes(path string) ([]*Node, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	out := []*Node{}
	inMath := false
	var mathLines []string
	inIf := false
	var ifLines []string
	var ifLeft, ifRight string

	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		if inMath {
			if line == "end" {
				out = append(out, &Node{Kind: NodeMathBlock, Lines: mathLines})
				inMath = false
				mathLines = nil
				continue
			}
			mathLines = append(mathLines, line)
			continue
		}
		if inIf {
			if line == "end" {
				childNodes := ParseLines(ifLines)
				n := &Node{Kind: NodeIf, Left: ifLeft, Right: ifRight, Children: childNodes}
				out = append(out, n)
				inIf = false
				ifLines = nil
				ifLeft = ""
				ifRight = ""
				continue
			}
			ifLines = append(ifLines, line)
			continue
		}
		if line == "math;" {
			inMath = true
			mathLines = []string{}
			continue
		}
		if strings.HasPrefix(line, "if ") || strings.HasPrefix(line, "\"if ") {
			cond := strings.TrimSpace(strings.TrimPrefix(line, "if"))
			cond = strings.TrimSpace(strings.TrimPrefix(cond, "\"if"))
			left := ""
			right := ""
			if idx := strings.Index(cond, "="); idx >= 0 {
				left = strings.TrimSpace(cond[:idx])
				right = strings.TrimSpace(cond[idx+1:])
				right = trimQuotes(right)
			}
			inIf = true
			ifLeft = left
			ifRight = right
			ifLines = []string{}
			continue
		}
		// single-line parse
		if strings.HasPrefix(line, "say ") {
			arg := strings.TrimSpace(line[len("say "):])
			out = append(out, &Node{Kind: NodeSay, Raw: raw, Expr: trimQuotes(arg)})
			continue
		}
		if strings.HasPrefix(line, "let ") {
			out = append(out, &Node{Kind: NodeAssign, Raw: raw, Expr: strings.TrimSpace(line[len("let "):])})
			continue
		}
		if strings.HasPrefix(line, "run ") {
			arg := trimQuotes(strings.TrimSpace(line[len("run "):]))
			out = append(out, &Node{Kind: NodeRun, Raw: raw, Args: []string{arg}})
			continue
		}
		if strings.HasPrefix(line, "build ") {
			parts := splitArgs(line)
			args := []string{}
			if len(parts) >= 2 {
				for _, p := range parts[1:] {
					args = append(args, trimQuotes(p))
				}
			}
			out = append(out, &Node{Kind: NodeBuild, Raw: raw, Args: args})
			continue
		}
		if strings.HasPrefix(line, "exec-safe ") || strings.HasPrefix(line, "exec ") {
			parts := splitArgs(line)
			if len(parts) >= 2 {
				out = append(out, &Node{Kind: NodeExec, Raw: raw, Args: parts[1:]})
			}
			continue
		}
		if line == "time.print" {
			out = append(out, &Node{Kind: NodeTimePrint, Raw: raw})
			continue
		}
		if strings.HasPrefix(line, "time.set") {
			parts := splitArgs(line)
			args := []string{}
			if len(parts) >= 2 {
				args = parts[1:]
			}
			out = append(out, &Node{Kind: NodeTimeSet, Raw: raw, Args: args})
			continue
		}
		// fallback math/expression
		out = append(out, &Node{Kind: NodeMath, Raw: raw, Expr: line})
	}
	return out, nil
}

func ParseLine(line string) ([]*Node, error) {
	return ParseLines([]string{line}), nil
}

func ParseLines(lines []string) []*Node {
	out := []*Node{}
	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "say ") {
			arg := strings.TrimSpace(line[len("say "):])
			out = append(out, &Node{Kind: NodeSay, Raw: raw, Expr: trimQuotes(arg)})
			continue
		}
		if strings.HasPrefix(line, "let ") {
			out = append(out, &Node{Kind: NodeAssign, Raw: raw, Expr: strings.TrimSpace(line[len("let "):])})
			continue
		}
		out = append(out, &Node{Kind: NodeMath, Raw: raw, Expr: line})
	}
	return out
}
