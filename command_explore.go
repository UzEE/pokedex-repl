package main

import "fmt"

func exploreCommand(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected area name in 1 argument, got %d", len(args))
	}

	name := args[0]

	printLine("Exploring %s...", name)

	area, err := c.client.GetLocationAreaDetails(name)
	if err != nil {
		return err
	}

	if len(area.PokemonEncounters) == 0 {
		printLine("No Pokémon encounters found.")
		return nil
	} else {
		printLine("Found Pokémon:")
	}

	for _, encounter := range area.PokemonEncounters {
		printLine(" - %s", encounter.Pokemon.Name)
	}

	c.currentArea = area

	return nil
}
