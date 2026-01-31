package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var colorMap = map[string]string{
	"green":  "\033[32m",
	"red":    "\033[31m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"reset":  "\033[0m",
}

func RunScript(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Cannot open file:", file)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var functions = make(map[string][]string)
	var currentFunc string
	var inFunc bool

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "func ") {
			currentFunc = strings.TrimSpace(strings.TrimPrefix(line, "func "))
			functions[currentFunc] = []string{}
			inFunc = true
			continue
		}

		if line == "end" && inFunc {
			currentFunc = ""
			inFunc = false
			continue
		}

		if inFunc {
			functions[currentFunc] = append(functions[currentFunc], line)
			continue
		}

		// Not in function → execute direct command
		executeLine(line, functions)
	}
}

func executeLine(line string, functions map[string][]string) {
	if strings.HasPrefix(line, "say(") && strings.HasSuffix(line, ")") {
		content := line[4 : len(line)-1]
		fmt.Println(content)
		return
	}

	if strings.HasPrefix(line, "text.color*") && strings.HasSuffix(line, "*") {
		colorName := line[11 : len(line)-1]
		if code, ok := colorMap[colorName]; ok {
			fmt.Print(code)
		}
		return
	}

	if fn, ok := functions[line]; ok {
		for _, l := range fn {
			executeLine(l, functions)
		}
		return
	}
}
