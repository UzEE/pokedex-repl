package main

import "fmt"

func exploreCommand(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected area name in 1 argument, got %d", len(args))
	}

	name := args[0]

	fmt.Printf("Exploring %s...\n", name)

	area, err := c.client.GetLocationAreaDetails(name)
	if err != nil {
		return err
	}

	if len(area.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon encounters found.")
		return nil
	} else {
		fmt.Println("Found Pokémon:")
	}

	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
