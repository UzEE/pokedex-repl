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
		fmt.Printf("Pokédex> ")
		input := readInput()

		err := handleCommand(input, commands, config)
		if err != nil {
			log.Println(err)
		}
	}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return sanitizeInput(scanner.Text())
}

func sanitizeInput(input string) string {
	output := strings.ToLower(input)
	return strings.TrimSpace(output)
}

func loadCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			handler:     helpCommand,
		},
		"map": {
			name:        "map",
			description: "Display the names of next 20 locations in the Pokémon world",
			handler:     mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of previous 20 locations in the Pokémon world",
			handler:     mapBCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokédex",
			handler:     exitCommand,
		},
		"exp": {
			name:        "exp",
			description: "Experiment command to test things out",
			handler:     expCommand,
		},
	}
}

func handleCommand(cmd string, commands map[string]command, config *config) error {
	if cmd == "" {
		return nil
	}

	c, ok := commands[cmd]
	if !ok {
		fmt.Printf("Command \"%s\" not found. Please type \"help\" to see a list of supported commands.\n", cmd)
		return nil
	}

	return c.handler(config)
}
