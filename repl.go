package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type config struct {
	pokeApiClient *pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	// Scan user input inifinite loop
	for {
		// Print the prompt
		fmt.Print("Pokedex > ")
		// Read the input
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		input := scanner.Text()
		// Clean the input
		cleanedInput := cleanInput(input)
		InputCommand := cleanedInput[0]
		arguments := cleanedInput[1:]

		// Check if the command is supported
		command, ok := getCommands()[InputCommand]

		if !ok {
			fmt.Printf("Unknown command\n")
			continue
		}

		// Execute the command
		err := command.callback(cfg, arguments...)

		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue
		}

	}
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	words := strings.Fields(trimmedText)

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show help",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 locations in the map",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations in the map",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show the caught pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
	}

}
