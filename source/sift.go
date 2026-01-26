package main

import (
	"strings"
)

func SplitWords(line string) []string {
	return strings.Fields(line)
}

func Contains(slice []string, word string) bool {
	for _, w := range slice {
		if w == word {
			return true
		}
	}
	return false
}

func Join(words []string, sep string) string {
	return strings.Join(words, sep)
}

func Index(slice []string, word string) int {
	for i, w := range slice {
		if w == word {
			return i
		}
	}
	return -1
}
