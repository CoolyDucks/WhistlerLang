package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// MIT License
//
// MIT License
// Copyright (c) 2026 CoolyDucks
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

func main() {
        fmt.Println("\033[36mWhistlerLang 1.1 , by CoolyDucks\033[0m")
	fmt.Println("Type 'quit' to exit, 'run <script.whlst>' to run a script, 'edit' for editor, 'help' for commands, 'source' for Open Source")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\033[32m>>> \033[0m")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		switch {
		case line == "quit":
			return
		case line == "help":
			printHelp()
		case line == "edit":
			runEditor()
		case strings.HasPrefix(line, "run "):
			file := strings.TrimSpace(line[4:])
			RunScript(file)
		case strings.HasPrefix(line, "say("):
			handleSay(line)
		case line == "source":
			installSource()
		default:
			fmt.Println("\033[31mUnknown:\033[0m", line)
		}
	}
}

func printHelp() {
	fmt.Println("\033[36mCommands:\033[0m")
	fmt.Println("  run <file>       -> run a WhistlerLang script")
	fmt.Println("  edit             -> open internal editor")
	fmt.Println("  quit             -> exit the interpreter")
	fmt.Println("  say(\"text\")     -> print green text")
	fmt.Println("  func <name>      -> define a function")
	fmt.Println("  math operations  -> add, sub, mul, div, mod")
	fmt.Println("  source           -> download Open Source repo")
}

func runEditor() {
	fmt.Println("\033[33mWhistlerLang 1.1 Editor (Official by CoolyDucks)\033[0m")
	fmt.Println("\033[33m~ Type code per line\033[0m")
	fmt.Println("\033[33mALT+Q to save and exit\033[0m")

	scanner := bufio.NewScanner(os.Stdin)
	var script []string
	for {
		fmt.Print("\033[33m~ \033[0m")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.EqualFold(line, "ALT+Q") {
			break
		}
		script = append(script, line)
	}

	fileName := "editor_output.whlst"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("\033[31mCannot create file:\033[0m", err)
		return
	}
	defer f.Close()
	for _, l := range script {
		f.WriteString(l + "\n")
	}
	fmt.Println("\033[32mSaved to\033[0m", fileName)
}

func handleSay(line string) {
	text := strings.TrimPrefix(line, "say(")
	text = strings.TrimSuffix(text, ")")
	text = strings.Trim(text, "\"")
	fmt.Println("\033[32m" + text + "\033[0m")
}

func installSource() {
	f, err := os.Create("installsource.sh")
	if err != nil {
		fmt.Println("\033[31mCannot create file:\033[0m", err)
		return
	}
	defer f.Close()

	content := `#!/bin/bash
git clone https://github.com/CoolyDucks/WhistlerLang
echo "WhistlerLang source installed successfully"
`
	f.WriteString(content)
	os.Chmod("installsource.sh", 0755)
	fmt.Println("\033[32minstallsource.sh created. Run ./installsource.sh to get source\033[0m")
}
