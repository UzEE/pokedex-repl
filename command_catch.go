package main

import "fmt"

func catchCommand(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("missing required argument: pokemon-name")
	}

	name := args[0]

	pokemon, err := c.client.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Println(pokemon)

	return nil
}
