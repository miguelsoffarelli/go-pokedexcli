package pokeapi

import (
	"io"
	"net/http"
)

const (
	BaseURL = "https://pokeapi.co/api/v2"
)

// Fetches url data from the api, stores it in cache and returns it
func (c *Client) FetchApiData(url *string) ([]byte, error) {
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(*url, data)

	return data, nil
}
