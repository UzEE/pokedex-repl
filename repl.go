package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func startRepl() {
	input := make(chan string)
	commands := loadCommands()

	for {
		fmt.Printf("Pokedex > ")
		go readInput(input)

		err := handleCommand(<-input, commands)
		if err != nil {
			log.Println(err)
		}
	}
}

func readInput(input chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	input <- scanner.Text()
}

func loadCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			handler:     helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			handler:     exitCommand,
		},
	}
}

func handleCommand(cmd string, commands map[string]command) error {
	if cmd == "" {
		return nil
	}

	c, ok := commands[cmd]
	if !ok {
		fmt.Printf("Command %s not found. Please type \"help\" to see a list of supported commands.\n", cmd)
		return nil
	}

	return c.handler()
}
