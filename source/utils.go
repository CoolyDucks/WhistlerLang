package main

import (
	"errors"
	"os"
	"strings"
)

func EnsureDirs() error {
	return os.MkdirAll("release", 0755)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func EnsureWhlstExtension(path string) error {
	if !strings.HasSuffix(path, ".whlst") {
		return errors.New("file must have .whlst extension")
	}
	return nil
}

func EnsureObjectExtension(path string) error {
	if !strings.HasSuffix(path, ".o") {
		return errors.New("file must have .o extension")
	}
	return nil
}

func SafeWriteFile(path string, data []byte) error {
	dir := ""
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			dir = path[:i]
			break
		}
	}
	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return os.WriteFile(path, data, 0644)
}

func ReadFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
