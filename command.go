package main

import (
	"fmt"

	"github.com/UzEE/pokedexcli/internal/api"
)

type config struct {
	client *api.Client

	Next     *string
	Previous *string
}

type command struct {
	name        string
	description string
	usage       string
	handler     func(c *config, args ...string) error
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
		"catch": {
			name:        "catch <pokemon name>",
			description: "Try to catch the specified Pokémon",
			usage:       "Format: catch <pokemon name>",
			handler:     catchCommand,
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
