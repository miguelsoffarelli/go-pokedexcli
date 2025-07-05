package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageURL *string) (Location, error) {
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		locationsResp := Location{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return Location{}, err
		}

		return locationsResp, nil
	}

	data, err := c.FetchApiData(&url)
	if err != nil {
		return Location{}, err
	}

	locationsResp := Location{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	return locationsResp, nil
}
