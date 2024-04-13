package api

import (
	"fmt"
	"io"

	"github.com/UzEE/pokedexcli/internal/api/types/location"
)

func (c *Client) GetRegionDetails(name string) (location.Region, error) {
	reqUrl := fmt.Sprintf("%s/region/%s", baseURL, name)

	cached, found := c.cache.Get(reqUrl)
	if found {
		region, err := unmarshalJSON[location.Region](cached)
		if err != nil {
			return location.Region{}, fmt.Errorf("failed to cast cache value to Region")
		}

		return region, nil
	}

	resp, err := c.client.Get(reqUrl)
	if err != nil {
		return location.Region{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return location.Region{}, fmt.Errorf("failed to fetch region data. Status Code: %d", resp.StatusCode)
	}

	if err != nil {
		return location.Region{}, err
	}

	c.cache.Add(reqUrl, body)

	region, err := unmarshalJSON[location.Region](body)

	return region, err
}
