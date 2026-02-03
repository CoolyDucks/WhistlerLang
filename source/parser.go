package main

import "strings"

type CommandType int

const (
	CMD_UNKNOWN CommandType = iota
	CMD_SAY
	CMD_DO_COLOR
	CMD_IF
	CMD_MATH
	CMD_MATH_BLOCK
	CMD_END
)

type Command struct {
	Type  CommandType
	Text  string
	Color string
	Expr  string
}

func ParseLine(line string) Command {
	line = strings.TrimSpace(line)
	switch {
	case strings.HasPrefix(line, `"say "`):
		return Command{Type: CMD_SAY, Text: strings.TrimPrefix(line, `"say `)}
	case strings.HasPrefix(line, ";do color "):
		return Command{Type: CMD_DO_COLOR, Color: strings.TrimPrefix(line, ";do color ")}
	case strings.HasPrefix(line, `"if say = "`):
		parts := strings.SplitN(line, `"`, 3)
		if len(parts) > 2 {
			return Command{Type: CMD_IF, Text: parts[2], Color: ""}
		}
	case line == "end":
		return Command{Type: CMD_END}
	case strings.HasPrefix(line, "math() "):
		return Command{Type: CMD_MATH, Expr: strings.TrimPrefix(line, "math() ")}
	case line == "math;":
		return Command{Type: CMD_MATH_BLOCK}
	}
	return Command{Type: CMD_UNKNOWN, Text: line}
}
