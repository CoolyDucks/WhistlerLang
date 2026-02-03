package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lastSay string
var vars = make(map[string]float64)
var colorMap = map[string]string{
	"reset":   "\033[0m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
}

func ExecuteLine(cmd Command) {
	switch cmd.Type {
	case CMD_SAY:
		lastSay = cmd.Text
		fmt.Println(lastSay)
	case CMD_DO_COLOR:
		code, ok := colorMap[cmd.Color]
		if !ok {
			code = colorMap["reset"]
		}
		fmt.Print(code)
		if lastSay != "" {
			fmt.Println(lastSay)
		}
		fmt.Print(colorMap["reset"])
	case CMD_MATH:
		res, err := EvalExpression(cmd.Expr, vars)
		if err != nil {
			fmt.Println("Math error:", err)
			return
		}
		fmt.Println(res)
	case CMD_IF:
		if lastSay == cmd.Text {
			ExecuteLine(Command{Type: CMD_DO_COLOR, Color: cmd.Color})
		}
	case CMD_END:
		return
	default:
		fmt.Println("Unknown command:", cmd.Text)
	}
}

func RunScript(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Cannot open file:", file)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	inMath := false
	var mathBlock []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		cmd := ParseLine(line)

		if inMath {
			if cmd.Type == CMD_END {
				for _, ml := range mathBlock {
					e := ParseLine(ml)
					ExecuteLine(e)
				}
				inMath = false
				mathBlock = mathBlock[:0]
			} else {
				mathBlock = append(mathBlock, line)
			}
			continue
		}

		if cmd.Type == CMD_MATH_BLOCK {
			inMath = true
			continue
		}

		ExecuteLine(cmd)
	}
}
