package main

import (
	"fmt"
	"io"
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

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println()
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println()
		return
	}
	defer file.Close()
	fmt.Println(string(content))
}
