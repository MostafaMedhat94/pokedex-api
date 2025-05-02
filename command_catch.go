package main

import (
	"fmt"
	"math/rand"
	"pokedex/internal/pokeapi"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("pokemon name is required")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeApiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// Create a new random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	if catchPokemon(pokemon, rng) {
		cfg.caughtPokemon[pokemon.Name] = pokemon

		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func catchPokemon(poke pokeapi.Pokemon, rng *rand.Rand) bool {
	const maxBaseExp = 300.0

	// Normalize base experience to a catch difficulty (0.0 ~ 1.0)
	catchDifficulty := float64(poke.BaseExperience) / maxBaseExp

	// Invert to make higher baseExperience harder to catch
	catchChance := 1.0 - catchDifficulty

	// Add a minimum chance floor (e.g. 10%)
	if catchChance < 0.1 {
		catchChance = 0.1
	}

	// Roll the dice!
	roll := rng.Float64()

	// fmt.Printf("Trying to catch %s...\n", poke.Name)
	fmt.Printf("Base Experience: %d\n", poke.BaseExperience)
	fmt.Printf("Catch chance: %.2f%%\n", catchChance*100)
	fmt.Printf("Rolled: %.2f\n", roll)

	return roll <= catchChance
}
