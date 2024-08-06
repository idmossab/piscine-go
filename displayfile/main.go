package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(arg) != 1 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := arg[0]

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println()
		return
	}
	fmt.Println(string(content))
}
