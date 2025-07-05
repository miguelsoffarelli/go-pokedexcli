package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListPokemon(name *string) (Pokemon, error) {
	url := BaseURL + "/pokemon/" + *name

	if data, ok := c.cache.Get(*name); ok {
		pokeResp := Pokemon{}
		err := json.Unmarshal(data, &pokeResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokeResp, nil
	}

	data, err := c.FetchApiData(&url)
	if err != nil {
		return Pokemon{}, err
	}

	pokeResp := Pokemon{}
	err = json.Unmarshal(data, &pokeResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokeResp, nil
}
