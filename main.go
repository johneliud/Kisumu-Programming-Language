package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/johneliud/Kisumu-Programming-Language/repl"
)

func main() {
	// Get the current user of the operating system
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get current user: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Hello %s! Welcome to Kisumu-Programming-Language\n\n", user.Username)
	fmt.Println("Proceed by typing in commands")
	// Starts the REPL
	repl.Start(os.Stdin, os.Stdout)
}
