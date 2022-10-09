package main

import "fmt"

type input struct {
	command string
}

func (input *input) ReturnCommand() string {
	return input.command
}

func main() {
	test := input{command: "Boo!"}

	fmt.Println(test.ReturnCommand())
}