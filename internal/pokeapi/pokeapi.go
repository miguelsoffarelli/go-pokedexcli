package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pokecache "github.com/miguelsoffarelli/go-pokedexcli/internal/pokecache"
)

func FetchLocationArea(url string, cache *pokecache.Cache) (Location, error) {
	var data Location
	// First check for the requested data in the cache
	rawData, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(rawData, &data)
		if err != nil {
			return Location{}, err
		}
		return data, nil
	}
	// If the requested data is not in the cache, proceed with the api call
	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Location{}, fmt.Errorf("response failed with status code %v", res.StatusCode)
	}
	if err != nil {
		return Location{}, err
	}
	// Save the retrieved data in the cache
	cache.Add(url, body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return Location{}, err
	}

	return data, nil
}

type Location struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
