package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(locationAreaName string) ([]PokemonShallow, error) {
	url := fmt.Sprintf("%s/%s/%s", c.baseURL, "location-area", locationAreaName)

	resp, err := c.fetchLocationPokemonEncounters(url)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch pokemon for location area %s: %w", locationAreaName, err)
	}
	if len(resp.PokemonEncounters) == 0 {
		return nil, fmt.Errorf("no pokemon found in %s", locationAreaName)
	}

	var pokemonList []PokemonShallow

	for _, encounter := range resp.PokemonEncounters {
		pokemonList = append(pokemonList, encounter.Pokemon)
	}

	return pokemonList, nil
}
func (c *Client) fetchLocationPokemonEncounters(url string) (LocationPokemonEncountersResponse, error) {
	cachedResponse, found := c.cache.Get(url)

	if found {
		var locationPokemonEncounters LocationPokemonEncountersResponse
		if err := json.Unmarshal(cachedResponse, &locationPokemonEncounters); err != nil {
			return LocationPokemonEncountersResponse{}, fmt.Errorf("failed to unmarshal cached response: %w", err)
		}
		return locationPokemonEncounters, nil
	}

	resp, err := c.httpClient.Get(url)

	if err != nil {
		return LocationPokemonEncountersResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return LocationPokemonEncountersResponse{}, fmt.Errorf("bad response (%d): %s", resp.StatusCode, string(body))
	}

	var locationPokemonEncounters LocationPokemonEncountersResponse

	if err := json.NewDecoder(resp.Body).Decode(&locationPokemonEncounters); err != nil {
		return LocationPokemonEncountersResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	jsonData, err := json.Marshal(locationPokemonEncounters)

	if err == nil {
		c.cache.Add(url, jsonData)
	}

	return locationPokemonEncounters, nil
}
