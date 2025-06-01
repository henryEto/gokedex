package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 locations",
		},
	}
}

func startRepl() {
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
			err := command.callback()
			if err != nil {
				fmt.Printf("%v error: %v", command.name, err.Error())
			}
		} else {

			fmt.Println("Unknown command")
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	words := strings.Fields(input)
	return words
}
