package main

import (
	"time"

	"git.eo13-dev.com/enordaz/gokedex/internal/pokeapi"
)

const prompt = "Pokedex > "

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeClient: pokeClient,
	}
	startRepl(cfg)
}
