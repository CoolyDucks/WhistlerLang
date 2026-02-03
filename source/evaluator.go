package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EvalExpression(expr string, vars map[string]float64) (float64, error) {
	expr = strings.ReplaceAll(expr, "ร—", "*")
	expr = strings.ReplaceAll(expr, "รท", "/")
	tokens := strings.Fields(expr)
	if len(tokens) != 3 {
		return 0, fmt.Errorf("invalid expression: %s", expr)
	}
	a, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, err
	}
	b, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return 0, err
	}
	switch tokens[1] {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", tokens[1])
	}
}
