package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, arg string) error {
	if len(*cfg.pokeDex) == 0 {
		return errors.New("you haven't caught any pokemon")
	}

	fmt.Println(" Your Pokedex:")
	for k := range *cfg.pokeDex {
		fmt.Printf("   - %v\n", k)
	}
	return nil
}
