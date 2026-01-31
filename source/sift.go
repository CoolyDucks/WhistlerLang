package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Sift(line string)[]string{
	line=strings.TrimSpace(line)
	re:=regexp.MustCompile(`[^\s]+`)
	return re.FindAllString(line,-1)
}

func runtimeSift(line string){
	parts:=Sift(line)
	fmt.Println("Sift result:",parts)
}
