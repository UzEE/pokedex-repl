package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/UzEE/pokedexcli/internal/api"
	"github.com/UzEE/pokedexcli/internal/api/types/pokemon"
	"github.com/UzEE/pokedexcli/internal/api/types/species"
)

func startRepl() {
	commands := loadCommands()
	client := api.NewClient()

	config := &config{
		client:  &client,
		pokedex: make(map[string]species.PokemonSpecies),
		box:     make([]pokemon.Pokemon, 0, 30),
	}

	for {
		fmt.Printf("\nPokédex> ")
		cmd, args := readInput()

		err := handleCommand(cmd, args, commands, config)
		if err != nil {
			log.Println(err)
		}
	}
}

func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return sanitizeInput(scanner.Text())
}

func sanitizeInput(input string) (command string, args []string) {
	output := strings.ToLower(input)
	output = strings.TrimSpace(output)
	parsed := strings.Fields(output)
	return parsed[0], parsed[1:]
}
