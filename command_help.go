package main

import "fmt"

func commandHelp(cfg *config, arg string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("  %v\t\t\t%v\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}
