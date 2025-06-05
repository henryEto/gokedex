package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	var locationsResp RespShallowLocations
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
