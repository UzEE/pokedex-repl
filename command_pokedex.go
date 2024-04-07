package main

import (
	"errors"
	"fmt"
)

func pokedexCommand(c *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("too many arguments")
	}

	count := len(c.pokedex)

	if count == 0 {
		fmt.Println("Your Pokédex is empty. Go catch some Pokémon!")
		return nil
	}

	fmt.Printf("Your have %d Pokémon registered in your Pokédex:\n", count)

	for _, pokemon := range c.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
