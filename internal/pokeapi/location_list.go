package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

var cache *pokecache.Cache = pokecache.NewCache(50 * time.Second)

// FetchNextLocations fetches the next 20 locations
func (c *Client) FetchNextLocations() (LocationAPIResponse, error) {
	if c.NextURL == "" {
		return LocationAPIResponse{}, fmt.Errorf("no more locations to fetch")
	}

	return c.fetchLocations(c.NextURL)
}

// FetchPreviousLocations fetches the previous 20 locations
func (c *Client) FetchPreviousLocations() (LocationAPIResponse, error) {
	if c.PrevURL == "" {
		return LocationAPIResponse{}, fmt.Errorf("already at the beginning")
	}

	return c.fetchLocations(c.PrevURL)
}

func (c *Client) fetchLocations(url string) (LocationAPIResponse, error) {
	if cachedResponse, err := checkCache(url); err == nil {
		c.NextURL = cachedResponse.Next
		c.PrevURL = cachedResponse.Previous

		fmt.Printf(">>>>>>>>>>>> Cache hit for URL: %s\n", url)

		return cachedResponse, nil
	}

	resp, err := c.httpClient.Get(url)

	if err != nil {
		return LocationAPIResponse{}, fmt.Errorf("failed to fetch locations: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return LocationAPIResponse{}, fmt.Errorf("bad response (%d): %s", resp.StatusCode, string(body))
	}

	var apiResp LocationAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return LocationAPIResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	if err := addToCache(url, apiResp); err != nil {
		formattedError := fmt.Errorf("failed to cache response: %w", err)

		fmt.Print(formattedError)
	}

	// Update next and previous URLs
	c.NextURL = apiResp.Next
	c.PrevURL = apiResp.Previous

	return apiResp, nil
}

func checkCache(url string) (LocationAPIResponse, error) {
	locationsAPIResponse := LocationAPIResponse{}

	if cachedResp, found := cache.Get(url); found {
		if err := json.Unmarshal(cachedResp, &locationsAPIResponse); err != nil {
			return locationsAPIResponse, fmt.Errorf("failed to decode cached response: %w", err)
		}

		return locationsAPIResponse, nil
	}

	return locationsAPIResponse, fmt.Errorf("cache miss")
}
func addToCache(url string, resp LocationAPIResponse) error {
	responseJSON, err := json.Marshal(resp)

	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	cache.Add(url, responseJSON)

	return nil
}
