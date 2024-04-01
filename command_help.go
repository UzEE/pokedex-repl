package main

import "fmt"

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
