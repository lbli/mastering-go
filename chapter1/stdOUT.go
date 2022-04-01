package main

import (
	"io"
	"os"
)

func main() {

	mystrings := ""

	arguments := os.Args

	if len(arguments) == 1 {
		mystrings = "please give me one arguments"
	} else {
		mystrings = arguments[1]
	}

	io.WriteString(os.Stdout, mystrings)
	io.WriteString(os.Stdout, "\n")

}
