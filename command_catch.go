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

	mon, err := c.client.GetPokemonAndSpecies(name)
	if err != nil {
		return err
	}

	formattedName := getFormattedName(c, mon)

	printLine("Throwing a Pokéball at %s...", formattedName)

	pokemonSpecies := mon.SpeciesEntry

	rate := pokemonSpecies.CaptureRate
	roll := rand.Int() % 255

	if roll < rate {
		printLine("%s was caught! (rate: %d, roll: %d)", formattedName, rate, roll)

		mon.SpeciesEntry = pokemonSpecies

		c.pokedex[pokemonSpecies.Name] = mon
		c.box = append(c.box, mon)

		printLine("You have registered %d Pokémon Species in your Pokédex.", len(c.pokedex))
		printLine("You have %d Pokémon in your box.", len(c.box))

		if c.currentEncounter != nil && c.currentEncounter.Name == mon.Name {
			c.currentEncounter = nil
		}

		c.latestCaught = &pokemonSpecies.Name

		printLine()
		printLine("Type:")
		printLine("  %sinspect%s to view the Pokédex entry on %s", c.colors.Blue, c.colors.Reset, formattedName)
		printLine("  %spokedex%s to view a list of all Pokémon in your Pokédex", c.colors.Blue, c.colors.Reset)

	} else {
		printLine("%s broke free! (rate: %d, roll: %d)", formattedName, rate, roll)
	}

	return nil
}
