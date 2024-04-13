package main

func helpCommand(_ *config, _ ...string) error {
	printLine()
	printLine("Welcome to the Pokedex!")
	printLine("Usage: <cmd> [args]")
	printLine()

	commands, command_order, err := loadCommands()
	if err != nil {
		return err
	}

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

	for _, name := range command_order {
		printLine(
			"%s: %*s %s %*s (%s)",
			name,
			maxName-getLen(name),
			"",
			commands[name].description,
			maxDesc-getLen(commands[name].description),
			"",
			commands[name].usage,
		)
	}

	printLine()
	return nil
}

func getLen(s string) int {
	return len([]rune(s))
}
