package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("please give me arguments more than one")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	//n,_:= strconv.ParseFloat(arguments[i],64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min", min)
	fmt.Println("Max", max)

}
