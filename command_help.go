package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	// Print the following
	/*
		Welcome to the Pokedex!
		Usage:

		help: Displays a help message
		exit: Exit the Pokedex
	*/
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map:  Show the next 20 locations in the map")
	fmt.Println("mapb: Show the previous 20 locations in the map")
	fmt.Println("explore: Explore a location's Pokemon encounters")
	fmt.Println("catch: Catch a pokemon")
	fmt.Println("inspect: Inspect a pokemon")
	fmt.Println("pokedex: Show the caught pokemon")

	return nil
}
