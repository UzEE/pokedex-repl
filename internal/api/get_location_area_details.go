package api

import (
	"fmt"
	"io"

	"github.com/UzEE/pokedexcli/internal/api/types/location"
)

func (c *Client) GetLocationAreaDetails(areaName string) (location.LocationArea, error) {
	reqUrl := fmt.Sprintf("%s/location-area/%s", baseURL, areaName)

	cached, found := c.cache.Get(reqUrl)
	if found {
		area, err := unmarshalJSON[location.LocationArea](cached)
		if err != nil {
			return location.LocationArea{}, fmt.Errorf("failed to cast cache value to LocationArea")
		}

		return area, nil
	}

	resp, err := c.client.Get(reqUrl)
	if err != nil {
		return location.LocationArea{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return location.LocationArea{}, fmt.Errorf("failed to fetch location area data. Status Code: %d", resp.StatusCode)
	}

	if err != nil {
		return location.LocationArea{}, err
	}

	c.cache.Add(reqUrl, body)

	area, err := unmarshalJSON[location.LocationArea](body)

	return area, err
}
