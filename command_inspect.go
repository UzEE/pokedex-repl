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

	fmt.Println("Name:", entry.Name)
	// print height, weight, abilities, types and stats
	fmt.Println("Height:", entry.Height)
	fmt.Println("Weight:", entry.Weight)

	fmt.Println("Abilities:")
	for _, ability := range entry.Abilities {
		fmt.Printf(" - %s\n", ability.Ability.Name)
	}

	fmt.Println("Types:")
	for _, t := range entry.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	fmt.Println("Stats:")
	for _, stat := range entry.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
