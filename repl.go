package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/UzEE/pokedexcli/internal/api"
	"github.com/UzEE/pokedexcli/internal/api/types/pokemon"
	"golang.org/x/term"
)

const LF = "\n"
const CRLF = "\r\n"

func startRepl() {
	commands := loadCommands()
	client := api.NewClient()

	config := &config{
		client:  &client,
		pokedex: make(Pokedex),
		box:     make([]pokemon.Pokemon, 0, 30),
	}

	if runtime.GOOS == "windows" {
		config.newline = LF
		makeStandardTerminal(commands, config)
	} else {
		config.newline = CRLF
		makeRawTerminal(commands, config)
	}
}

func makeRawTerminal(commands map[string]command, config *config) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	terminal := term.NewTerminal(os.Stdin, "\r\nPokédex> ")

	for {
		line, err := terminal.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		cmd, args := sanitizeInput(line)

		err = handleCommand(cmd, args, commands, config)
		if err != nil {
			log.Println(err)
		}
	}
}

func makeStandardTerminal(commands map[string]command, config *config) {
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
