package pokeapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	data, ok := c.cache.Get(url)
	if !ok {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}
		c.cache.Add(url, data)
	}

	var locationsResp RespShallowLocations
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
