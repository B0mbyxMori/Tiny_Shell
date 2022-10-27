package main

import "testing"

func Test_ReturnCommand(t *testing.T) {
	testInput := &input{command: "Hey."}

	if testInput.command != "Hey." {
		t.Error("Incorrect string: Expected 'Hey.', got ", testInput)
	}
}