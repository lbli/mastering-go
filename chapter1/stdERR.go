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

	io.WriteString(os.Stdout, "this is standout output\n")
	io.WriteString(os.Stderr, mystrings)
	io.WriteString(os.Stderr, "\n")
}
