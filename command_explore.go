package main

import (
	"fmt"
)

func commandExplore(cfg *config, areaName string) error {
	encounters, err := cfg.pokeClient.ListEncounters(&areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", encounters.Name)
	fmt.Println("Found Pokemon:")
	for _, p := range encounters.Encounters {
		fmt.Printf("  - %s\n", p.Pokemon.Name)
	}

	return nil
}
