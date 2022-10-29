// Responsibilities: Handles all input from user.
// Depends on user_interface. Relays to os_interface.

package main

import "fmt"

type input struct {
	command string
}

// NOTE: Move to output_handler eventually.
func (input *input) ReturnCommand() string {
	return input.command
}

// NOTE: Will I need this?
// func (input *input) ParseCommand() []string {

// }

