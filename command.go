package main

import (
	"fmt"

	"github.com/UzEE/pokedexcli/internal/api"
	"github.com/UzEE/pokedexcli/internal/api/types/location"
	"github.com/UzEE/pokedexcli/internal/api/types/pokemon"
	"golang.org/x/term"
)

type Pokedex map[string]pokemon.Pokemon

type config struct {
	client *api.Client

	prompt string
	colors term.EscapeCodes

	NextRegion       *string
	PreviousRegion   *string
	LocationPage     int
	LocationPageSize int
	NextArea         *string
	PreviousArea     *string

	currentRegion    *location.Region
	currentLocation  *location.Location
	currentArea      *location.LocationArea
	currentEncounter *pokemon.Pokemon

	latestCaught *string
	pokedex      Pokedex
	box          []pokemon.Pokemon
}

type command struct {
	name        string
	description string
	usage       string
	handler     func(c *config, args ...string) error
}

func loadCommands() (map[string]command, []string, error) {
	commands := map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			usage:       "Format: help",
			handler:     helpCommand,
		},
		"region": {
			name:        "region",
			description: "Display the names of next 20 regions in the Pokémon world",
			usage:       "Format: region",
			handler:     regionCommand,
		},
		"visit": {
			name:        "visit",
			description: "Visit the specified region from region command and view the locations in that region",
			usage:       "Format: visit <region_name>",
			handler:     visitCommand,
		},
		"explore": {
			name:        "explore",
			description: "Explore the specified location from visit command and view the areas in that location",
			usage:       "Format: explore <location_name>",
			handler:     exploreCommand,
		},
		"map": {
			name:        "map [location_name]",
			description: "Display the names of areas in the specified location. If no location_name is given, then display the names of next 20 locations in the Pokémon world",
			usage:       "Format: map [location_name] or map",
			handler:     mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of previous 20 locations in the Pokémon world",
			usage:       "Format: mapb",
			handler:     mapBCommand,
		},
		"scan": {
			name:        "scan",
			description: "PokeScan the specified area for wild Pokémon. Use the map command to see the areas in the location.",
			usage:       "Format: scan <area_name>",
			handler:     scanCommand,
		},
		"encounter": {
			name:        "encounter",
			description: "Encounter a wild Pokémon",
			usage:       "Format: encounter",
			handler:     encounterCommand,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try to catch the specified Pokémon",
			usage:       "Format: catch <pokemon_name>",
			handler:     catchCommand,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View the Pokédex entry for the specified Pokémon",
			usage:       "Format: inspect <pokemon_name>",
			handler:     inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View names of all the Pokémon in your Pokédex",
			usage:       "Format: pokedex",
			handler:     pokedexCommand,
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

	command_order := []string{
		"region",
		"visit",
		"explore",
		"map",
		"mapb",
		"scan",
		"encounter",
		"catch",
		"inspect",
		"pokedex",
		"help",
		"exit",
		"exp",
	}

	if len(commands) != len(command_order) {
		return nil, nil, fmt.Errorf("command count mismatch")
	}

	return commands, command_order, nil
}

func handleCommand(cmd string, args []string, commands map[string]command, config *config) error {
	if cmd == "" {
		return nil
	}

	if config.currentEncounter != nil {
		switch cmd {
		case "throw":
			return catchCommand(config, config.currentEncounter.Name)
		case "run":
			return runCommand(config)
		default:
			printLine("You are in an encounter with %s%s%s. Please type %s\"throw\"%s to throw a Pokéball or %s\"run\"%s to run away.", config.colors.Green, config.currentEncounter.Name, config.colors.Reset, config.colors.Blue, config.colors.Reset, config.colors.Blue, config.colors.Reset)
			return nil
		}
	}

	c, ok := commands[cmd]
	if !ok {
		printLine("Command \"%s\" not found. Please type \"help\" to see a list of supported commands.", cmd)
		return nil
	}

	return c.handler(config, args...)
}
