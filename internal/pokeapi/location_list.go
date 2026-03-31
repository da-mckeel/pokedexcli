package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(url *string) (Locations, error) {
	pageURL := baseURL + "/location-area"
	if url != nil {
		pageURL = *url
	}

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	var locations Locations
	if err := json.Unmarshal(data, &locations); err != nil {
		return Locations{}, err
	}

	return locations, nil
}
