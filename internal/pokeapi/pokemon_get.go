package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("%s/%s/%s", c.baseURL, "pokemon", name)

	cachedResponse, found := c.cache.Get(url)

	if found {
		var pokemon Pokemon
		if err := json.Unmarshal(cachedResponse, &pokemon); err != nil {
			return Pokemon{}, fmt.Errorf("failed to unmarshal cached response: %w", err)
		}
		return pokemon, nil
	}

	resp, err := c.httpClient.Get(url)

	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to fetch pokemon: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return Pokemon{}, fmt.Errorf("pokemon not found: %s", name)
		}

		body, _ := io.ReadAll(resp.Body)
		return Pokemon{}, fmt.Errorf("bad response (%d): %s", resp.StatusCode, string(body))
	}

	var pokemon Pokemon

	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("failed to decode response: %w", err)
	}

	jsonData, err := json.Marshal(pokemon)

	if err == nil {
		c.cache.Add(url, jsonData)
	}

	return pokemon, nil
}
