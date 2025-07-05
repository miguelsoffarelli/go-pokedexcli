package pokeapi

import (
	"net/http"
	"time"

	"github.com/miguelsoffarelli/go-pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	Pokedex    map[string]Pokemon
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		Pokedex: make(map[string]Pokemon),
	}
}
