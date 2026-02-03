package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("WhistlerLang 1.2 , by CoolyDucks")
	fmt.Println("Type 'help' to see commands, 'quit' to exit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")
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
		default:
			cmd := ParseLine(line)
			ExecuteLine(cmd)
		}
	}
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  run <file>")
	fmt.Println("  edit")
	fmt.Println("  quit")
	fmt.Println("  say <text>")
	fmt.Println("  if say = <text>  -> conditional color")
	fmt.Println("  do color <color>")
	fmt.Println("  math() <expr>")
	fmt.Println("  math; ... end")
}

func runEditor() {
	fmt.Println("Editor Mode, type code per line, ALT+Q to save and exit")
	scanner := bufio.NewScanner(os.Stdin)
	var script []string
	for {
		fmt.Print("~ ")
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
		fmt.Println("Cannot create file:", err)
		return
	}
	defer f.Close()
	for _, l := range script {
		f.WriteString(l + "\n")
	}
	fmt.Println("Saved to", fileName)
}
