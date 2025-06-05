package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"git.eo13-dev.com/enordaz/gokedex/internal/pokeapi"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 locations",
			callback:    commandMapb,
		},
	}
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		inputCommand := words[0]
		command, exists := getCommands()[inputCommand]

		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Printf("%v error: %v\n", command.name, err.Error())
			}
		} else {

			fmt.Println("unknown command")
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeClient       pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	words := strings.Fields(input)
	return words
}
