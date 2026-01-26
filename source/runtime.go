package main

import (
	"fmt"
	"os"
)

func LoadScript(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Cannot read file:", file)
		return nil
	}
	lines := make([]string, 0)
	for _, l := range SplitLines(string(data)) {
		if l != "" {
			lines = append(lines, l)
		}
	}
	return lines
}

func SplitLines(data string) []string {
	lines := []string{}
	start := 0
	for i, c := range data {
		if c == '\n' || c == '\r' {
			if start < i {
				lines = append(lines, data[start:i])
			}
			start = i + 1
		}
	}
	if start < len(data) {
		lines = append(lines, data[start:])
	}
	return lines
}
