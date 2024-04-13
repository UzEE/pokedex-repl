package main

import (
	"fmt"
	"math/rand"
)

func encounterCommand(c *config, _ ...string) error {
	area := c.currentArea

	if area == nil {
		return fmt.Errorf("you are not exploring any area")
	}

	printLine("Looking for Pokémon in %s...", area.Name)

	roll := rand.Int() % 255

	possiblePokemon := make([]string, 0, len(area.PokemonEncounters))

	for _, encounter := range area.PokemonEncounters {
		maxChance := 0

		for _, version := range encounter.VersionDetails {
			if version.MaxChance > maxChance {
				maxChance = version.MaxChance
			}
		}

		if roll < maxChance {
			possiblePokemon = append(possiblePokemon, encounter.Pokemon.Name)
		}
	}

	if len(possiblePokemon) == 0 {
		printLine("No Pokémon found.")
		return nil
	}

	pokemon := possiblePokemon[rand.Int()%len(possiblePokemon)]

	details, err := c.client.GetPokemonAndSpecies(pokemon)
	if err != nil {
		return err
	}

	c.currentEncounter = &details
	formattedName := getFormattedName(c, details)

	printLine("A wild %s appeared!", formattedName)

	printLine()
	printLine("Type:")
	printLine("  %sthrow%s a Pokéball at %s", c.colors.Blue, c.colors.Reset, formattedName)
	printLine("  %srun%s to get away", c.colors.Blue, c.colors.Reset)

	return nil
}
