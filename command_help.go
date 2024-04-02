package main

import "fmt"

func helpCommand(_ *config, _ ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: <cmd> [args]")
	fmt.Println()

	commands := loadCommands()

	maxName, maxDesc := 0, 0
	for name, cmd := range commands {
		if len(name) > maxName {
			maxName = len(name)
		}

		if len(cmd.description) > maxDesc {
			maxDesc = len(cmd.description)
		}
	}

	maxName += 4
	maxDesc += 4

	for name, cmd := range commands {
		fmt.Printf("%s: %*s %s %*s (%s)\n", name, maxName-len(name), "", cmd.description, maxDesc-len(cmd.description), " ", cmd.usage)
	}

	fmt.Println()
	return nil
}
