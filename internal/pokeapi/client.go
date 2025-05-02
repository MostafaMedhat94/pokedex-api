package pokeapi

import (
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type Client struct {
	cache      *pokecache.Cache
	httpClient *http.Client
	baseURL    string

	NextURL string
	PrevURL string
}

func NewClient(timeout, cacheTTL time.Duration) *Client {
	return &Client{
		cache:      pokecache.NewCache(cacheTTL),
		httpClient: &http.Client{Timeout: timeout},
		baseURL:    "https://pokeapi.co/api/v2",
		NextURL:    "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		PrevURL:    "",
	}
}
