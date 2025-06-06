package pokeapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(areaName *string) (RespShallowLocationExplore, error) {
	if areaName == nil {
		return RespShallowLocationExplore{}, errors.New("invalid location for exploration")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + *areaName
	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocationExplore{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocationExplore{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocationExplore{}, err
		}
		c.cache.Add(url, data)
	}

	var encounterResp RespShallowLocationExplore

	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&encounterResp); err != nil {
		return RespShallowLocationExplore{}, err
	}

	return encounterResp, nil
}

func (c *Client) PokemonInfo(pokemon *string) (Pokemon, error) {
	if pokemon == nil {
		return Pokemon{}, errors.New("please enter the pokemon name")
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + *pokemon
	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, data)
	}

	var pokemonResp Pokemon

	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&pokemonResp); err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}
