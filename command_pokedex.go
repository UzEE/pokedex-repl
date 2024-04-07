package main

import (
	"errors"
)

func pokedexCommand(c *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("too many arguments")
	}

	count := len(c.pokedex)

	if count == 0 {
		printLine("Your Pokédex is empty. Go catch some Pokémon!")
		return nil
	}

	printLine("Your have %d Pokémon registered in your Pokédex:", count)

	for _, pokemon := range c.pokedex {
		printLine(" - %s", pokemon.Name)
	}

	return nil
}
