package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	data, ok := c.cache.Get(url)
	if ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bas status code: %v", res.StatusCode)
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemonResp, nil
}
