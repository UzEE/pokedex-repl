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
		if getLen(name) > maxName {
			maxName = getLen(name)
		}

		if getLen(cmd.description) > maxDesc {
			maxDesc = getLen(cmd.description)
		}
	}

	maxName += 4
	maxDesc += 4

	for name, cmd := range commands {
		fmt.Printf(
			"%s: %*s %s %*s (%s)\n",
			name,
			maxName-getLen(name),
			"",
			cmd.description,
			maxDesc-getLen(cmd.description),
			"",
			cmd.usage,
		)
	}

	fmt.Println()
	return nil
}

func getLen(s string) int {
	return len([]rune(s))
}
