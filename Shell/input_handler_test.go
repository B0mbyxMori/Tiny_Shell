// Responsibilities: Contains all input_handler tests.

package main

import "testing"

func TestReturnCommand(t *testing.T) {
	testInput := &input{command: "Hey."}

	if testInput.command != "Hey." {
		t.Error("Incorrect string: Expected 'Hey.', got ", testInput)
	}
}