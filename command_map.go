package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
)

func commandMapForward(cfg *config, args ...string) error {
	locations, err := cfg.pokeApiClient.FetchNextLocations()

	if err != nil {
		return err
	}

	printLocations(locations)

	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	locations, err := cfg.pokeApiClient.FetchPreviousLocations()

	if err != nil {
		return err
	}

	printLocations(locations)

	return nil
}

func printLocations(locationsResponse pokeapi.LocationAPIResponse) {
	for _, location := range locationsResponse.Results {
		fmt.Printf("%s\n", location.Name)
	}
}
