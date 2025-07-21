package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, ok := (*cfg.pokeDex)[pokemonName]
	if !ok {
		return errors.New("you havent caught that pokemon yet")
	}
	fmt.Println("")
	fmt.Printf(" Name: %v\n", pokemon.Name)
	fmt.Printf(" Height: %v\n", pokemon.Height)
	fmt.Printf(" Weight: %v\n", pokemon.Weight)
	fmt.Println(" Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("   -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println(" Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("   - %v\n", t.Type.Name)
	}
	fmt.Println("")

	return nil
}
