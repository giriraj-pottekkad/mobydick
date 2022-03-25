package main

import (
	handler "mobydick/pkg/handler"
	"os"
)

const (
	defultFilePath   = "mobydick.txt"
	defaultWordCount = 20
)

func main() {

	var messageSrc string
	if len(os.Args) <= 1 {
		//File Path not provided in arguments. Using default File
		messageSrc = defultFilePath

	} else {
		messageSrc = os.Args[1]
	}

	handler.ReadFile(messageSrc, defaultWordCount)
}
