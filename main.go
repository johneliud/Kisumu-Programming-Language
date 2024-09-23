package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/johneliud/Kisumu-Programming-Language/repl"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}

	// Get the current user of the operating system
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get current user: %s\n", err)
	}
	
	fmt.Printf("Hello %s! Welcome to Kisumu-Programming-Language\n\n", user.Username)
	repl.Start(os.Stdout)
}
