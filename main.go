package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		pokeApiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
