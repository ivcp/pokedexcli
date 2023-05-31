package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseUrl + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bas status code: %v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationsAreasResp := LocationAreasResp{}

	err = json.Unmarshal(data, &locationsAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationsAreasResp, nil
}
