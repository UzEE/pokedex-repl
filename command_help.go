package main

func helpCommand(_ *config, _ ...string) error {
	printLine()
	printLine("Welcome to the Pokedex!")
	printLine("Usage: <cmd> [args]")
	printLine()

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
		printLine(
			"%s: %*s %s %*s (%s)",
			name,
			maxName-getLen(name),
			"",
			cmd.description,
			maxDesc-getLen(cmd.description),
			"",
			cmd.usage,
		)
	}

	printLine()
	return nil
}

func getLen(s string) int {
	return len([]rune(s))
}
