package api

import (
	"fmt"
	"io"

	"github.com/UzEE/pokedexcli/internal/api/types/location"
)

func (c *Client) GetLocationDetails(locationName string) (location.Location, error) {
	reqUrl := fmt.Sprintf("%s/location/%s", baseURL, locationName)

	cached, found := c.cache.Get(reqUrl)
	if found {
		loc, err := unmarshalJSON[location.Location](cached)
		if err != nil {
			return location.Location{}, fmt.Errorf("failed to cast cache value to Location")
		}

		return loc, nil
	}

	resp, err := c.client.Get(reqUrl)
	if err != nil {
		return location.Location{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return location.Location{}, fmt.Errorf("failed to fetch location data. Status Code: %d", resp.StatusCode)
	}

	if err != nil {
		return location.Location{}, err
	}

	c.cache.Add(reqUrl, body)

	loc, err := unmarshalJSON[location.Location](body)

	return loc, err
}
