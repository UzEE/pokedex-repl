package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/UzEE/pokedexcli/internal/api"
)

func startRepl() {
	commands := loadCommands()
	client := api.NewClient()

	config := &config{
		client: &client,
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

func loadCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			usage:       "Format: help",
			handler:     helpCommand,
		},
		"map": {
			name:        "map",
			description: "Display the names of next 20 locations in the Pokémon world",
			usage:       "Format: map",
			handler:     mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of previous 20 locations in the Pokémon world",
			usage:       "Format: mapb",
			handler:     mapBCommand,
		},
		"explore": {
			name:        "explore <area name>",
			description: "Explore the specified area from map command",
			usage:       "Format: explore <area name>",
			handler:     exploreCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokédex",
			usage:       "Format: exit",
			handler:     exitCommand,
		},
		"exp": {
			name:        "exp",
			description: "Experiment command to test things out",
			usage:       "Format: exp",
			handler:     expCommand,
		},
	}
}

func handleCommand(cmd string, args []string, commands map[string]command, config *config) error {
	if cmd == "" {
		return nil
	}

	c, ok := commands[cmd]
	if !ok {
		fmt.Printf("Command \"%s\" not found. Please type \"help\" to see a list of supported commands.\n", cmd)
		return nil
	}

	return c.handler(config, args...)
}
