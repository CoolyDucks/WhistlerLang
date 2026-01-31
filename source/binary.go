package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func BinaryRun(app,command string){
	args:=strings.Fields(command)
	cmd:=exec.Command(app,args...)
	cmd.Stdout=cmd.Stdout
	cmd.Stderr=cmd.Stderr
	err:=cmd.Run()
	if err!=nil{
		fmt.Println("Error running",app,"with command",command,":",err)
	}
}

func runtimeBinary(line string){
	if strings.HasPrefix(line,"using("){
		app:=strings.TrimSuffix(strings.TrimPrefix(line,"using("),")")
		command:=""
		for _,part:=range strings.Split(line,"\n"){
			if strings.HasPrefix(part,"COOMAND="){
				command=strings.TrimPrefix(part,"COOMAND=")
			}
		}
		BinaryRun(app,command)
	}
}
