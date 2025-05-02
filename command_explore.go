package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("location area name is required")
	}

	locationAreaName := args[0]

	pokemonList, err := cfg.pokeApiClient.ListPokemon(locationAreaName)

	if err != nil {
		return err
	}
	if len(pokemonList) == 0 {
		fmt.Printf("No pokemon found in %s\n", locationAreaName)
		return nil
	}

	printPokemonList(pokemonList)

	return nil
}

func printPokemonList(pokemonList []pokeapi.PokemonShallow) {
	for _, pokemon := range pokemonList {
		fmt.Printf("%s\n", pokemon.Name)
	}
}
