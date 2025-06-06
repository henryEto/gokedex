package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemonName string) error {
	pokemon, err := cfg.pokeClient.PokemonInfo(&pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	baseExp := pokemon.BaseExp
	catchRate := 500 - baseExp
	catchRoll := rand.Intn(500)
	if catchRoll <= catchRate {
		(*cfg.pokeDex)[pokemonName] = pokemon
		fmt.Println(pokemonName, "was caught!")
	} else {
		fmt.Println(pokemonName, "escaped!")
	}

	return nil
}
