package main

import "fmt"

func inspectCommand(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("missing required argument: <pokemon name>")
	}

	name := args[0]

	entry, ok := c.pokedex[name]
	if !ok {
		return fmt.Errorf("you haven't caught %s", name)
	}

	printLine("Name:", entry.Name)
	// print height, weight, abilities, types and stats
	printLine("Height:", entry.Height)
	printLine("Weight:", entry.Weight)

	printLine("Abilities:")
	for _, ability := range entry.Abilities {
		printLine(" - %s", ability.Ability.Name)
	}

	printLine("Types:")
	for _, t := range entry.Types {
		printLine(" - %s", t.Type.Name)
	}

	printLine("Stats:")
	for _, stat := range entry.Stats {
		printLine(" - %s: %d", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
