package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const Licence = `MIT License
Copyright (c) 2026 CoolyDucks
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software are
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.`

func RunWap(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Cannot read WAP:", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			Exec(line)
		}
	}
}

func Exec(line string) {
	if strings.HasPrefix(line, "say(") && strings.HasSuffix(line, ")") {
		Say(line[4 : len(line)-1])
	} else if strings.HasPrefix(line, "build ") {
		target := strings.TrimSpace(line[6:])
		BuildWap(target)
	} else if line == "licence" || line == "license" {
		fmt.Println(Licence)
	} else {
		fmt.Println("Unknown:", line)
	}
}

func Say(msg string) {
	fmt.Println(msg)
}

func BuildWap(folder string) {
	output := "build-whistlerlang-app.wap"
	cmd := exec.Command("tar", "-czf", output, folder)
	cmd.Run()
	fmt.Println("Created WAP:", output)
}

func main() {
	fmt.Println("\033[32mWhistlerLang 1.0, by CoolyDucks")
	fmt.Println("Type 'quit' to exit, 'run <script.whlst>' to run a script, 'runwap <file.wap>' for WAP, 'help' for commands\033[0m")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\033[32m>>>\033[0m ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "quit" {
			break
		} else if strings.HasPrefix(line, "run ") {
			file := strings.TrimSpace(line[4:])
			RunWap(file)
		} else if strings.HasPrefix(line, "runwap ") {
			file := strings.TrimSpace(line[7:])
			RunWap(file)
		} else if line != "" {
			Exec(line)
		}
	}
}
