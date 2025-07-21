package main

import (
	"time"

	"git.eo13-dev.com/enordaz/gokedex/internal/pokeapi"
)

const prompt = "Pokedex > "

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	catchedPokemon := map[string]pokeapi.Pokemon{}
	cfg := &config{
		pokeClient: pokeClient,
		pokeDex:    &catchedPokemon,
	}
	startRepl(cfg)
}
