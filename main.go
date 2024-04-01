package main

import (
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	handler     func() error
}

func main() {
	startRepl()
}

func helpCommand() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := loadCommands()

	for name, cmd := range commands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}

	fmt.Println()
	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}
