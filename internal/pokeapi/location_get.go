package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (LocationResp, error) {
	endpoint := "/location-area/"
	fullURL := baseUrl + endpoint + location

	data, ok := c.cache.Get(fullURL)
	if ok {
		locationsResp := LocationResp{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return LocationResp{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationResp{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResp{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationResp{}, fmt.Errorf("bas status code: %v", res.StatusCode)
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationResp{}, err
	}

	locationsResp := LocationResp{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationsResp, nil
}
