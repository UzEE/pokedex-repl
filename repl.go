package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/UzEE/pokedexcli/internal/api"
	"github.com/UzEE/pokedexcli/internal/api/types/pokemon"
	"golang.org/x/term"
)

var ErrExitSafe = errors.New("exit safe")

func startRepl() {
	commands, _, err := loadCommands()
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewClient()

	config := &config{
		client:  &client,
		prompt:  "Pokédex> ",
		pokedex: make(Pokedex),
		box:     make([]pokemon.Pokemon, 0, 30),

		LocationPage:     0,
		LocationPageSize: 20,
	}

	if runtime.GOOS == "windows" {
		makeStandardTerminal(commands, config)
	} else {
		makeRawTerminal(commands, config)
	}

	os.Exit(0)
}

func makeRawTerminal(commands map[string]command, config *config) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	// we need to restore the terminal to its original state (cooked mode) before exiting
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	terminal := term.NewTerminal(os.Stdin, "\r\n"+config.prompt)
	config.colors = *terminal.Escape

	for {
		terminal.SetPrompt("\r\n" + config.prompt)
		line, err := terminal.ReadLine()

		if err != nil {

			// handle Ctrl+C and Ctrl+D gracefully
			if err == io.EOF {
				printLine("Goodbye!")
			} else {
				fmt.Printf("\r\n%v", err)
			}

			return
		}

		if line == "" {
			continue
		}

		cmd, args := sanitizeInput(line)

		err = handleCommand(cmd, args, commands, config)
		if err != nil {
			if err == ErrExitSafe {
				return
			}

			log.Println(err)
		}
	}
}

func makeStandardTerminal(commands map[string]command, config *config) {
	for {
		fmt.Printf("\n%s", config.prompt)
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

func printLine(args ...any) {
	if len(args) == 0 {
		fmt.Printf("\r\n")
		return
	}

	parts := make([]any, 0, len(args))
	for _, val := range args {
		str, ok := val.(string)
		if ok {
			str = strings.ReplaceAll(str, "\n", "\r\n")
			parts = append(parts, str)
		} else {
			parts = append(parts, val)
		}
	}

	fmt.Printf(parts[0].(string), parts[1:]...)
	fmt.Printf("\r\n")
}

func stripBreaks(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.ReplaceAll(s, "\f", " ")

	return s
}
