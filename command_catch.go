package main

import (
	"fmt"
	"math/rand"
)

func catchCommand(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("missing required argument: <pokemon name>")
	}

	name := args[0]

	mon, err := c.client.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokéball at %s...\n", mon.Name)

	pokemonSpecies, err := c.client.GetPokemonSpecies(mon.Species.Name)
	if err != nil {
		return err
	}

	rate := pokemonSpecies.CaptureRate
	roll := rand.Int() % 255

	if roll < rate {
		fmt.Printf("%s was caught!\n", mon.Name)

		mon.SpeciesEntry = pokemonSpecies

		c.pokedex[pokemonSpecies.Name] = mon
		c.box = append(c.box, mon)

		fmt.Printf("You have registered %d Pokémon Species in your Pokédex.\n", len(c.pokedex))
		fmt.Printf("You have %d Pokémon in your box.\n", len(c.box))

	} else {
		fmt.Printf("%s broke free! (rate: %d, roll: %d)", mon.Name, rate, roll)
	}

	return nil
}
