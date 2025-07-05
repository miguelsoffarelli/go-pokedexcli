package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
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

	req, err := http.NewRequest("GET", *areaURL, nil)
	if err != nil {
		return Area{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Area{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Area{}, err
	}

	areaResp := Area{}
	err = json.Unmarshal(data, &areaResp)
	if err != nil {
		return Area{}, err
	}

	c.cache.Add(*areaURL, data)
	return areaResp, nil
}
