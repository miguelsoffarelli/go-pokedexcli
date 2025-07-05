package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListArea(areaURL *string) (Area, error) {
	if data, ok := c.cache.Get(*areaURL); ok {
		areaResp := Area{}
		err := json.Unmarshal(data, &areaResp)
		if err != nil {
			return Area{}, err
		}

		return areaResp, nil
	}

	data, err := c.FetchApiData(areaURL)
	if err != nil {
		return Area{}, err
	}

	areaResp := Area{}
	err = json.Unmarshal(data, &areaResp)
	if err != nil {
		return Area{}, err
	}

	return areaResp, nil
}
