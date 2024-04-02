package api

import (
	"fmt"
	"io"
)

func (c *Client) GetLocationAreaDetails(areaName string) (LocationArea, error) {
	reqUrl := fmt.Sprintf("%s/location-area/%s", baseURL, areaName)

	cached, found := c.cache.Get(reqUrl)
	if found {
		area, err := unmarshalJSON[LocationArea](cached)
		if err != nil {
			return LocationArea{}, fmt.Errorf("failed to cast cache value to LocationArea")
		}

		return area, nil
	}

	resp, err := c.client.Get(reqUrl)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return LocationArea{}, fmt.Errorf("failed to fetch location area data. Status Code: %d", resp.StatusCode)
	}

	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(reqUrl, body)

	area, err := unmarshalJSON[LocationArea](body)

	return area, err
}
