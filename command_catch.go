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

	printLine("Throwing a Pokéball at %s...", mon.Name)

	pokemonSpecies, err := c.client.GetPokemonSpecies(mon.Species.Name)
	if err != nil {
		return err
	}

	rate := pokemonSpecies.CaptureRate
	roll := rand.Int() % 255

	if roll < rate {
		printLine("%s was caught!", mon.Name)

		mon.SpeciesEntry = pokemonSpecies

		c.pokedex[pokemonSpecies.Name] = mon
		c.box = append(c.box, mon)

		printLine("You have registered %d Pokémon Species in your Pokédex.", len(c.pokedex))
		printLine("You have %d Pokémon in your box.", len(c.box))

	} else {
		printLine("%s broke free! (rate: %d, roll: %d)", mon.Name, rate, roll)
	}

	return nil
}
