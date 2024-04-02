package api

import (
	"fmt"
	"io"
)

func (c *Client) ListLocationArea(url *string) (PagedResourceList, error) {
	reqUrl := "https://pokeapi.co/api/v2/location-area"

	if url != nil {
		reqUrl = *url
	}

	cached, found := c.cache.Get(reqUrl)
	if found {
		list, err := unmarshalJSON[PagedResourceList](cached)
		if err != nil {
			return PagedResourceList{}, fmt.Errorf("failed to cast cache value to PagedResourceList")
		}

		return list, nil
	}

	resp, err := c.client.Get(reqUrl)
	if err != nil {
		return PagedResourceList{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return PagedResourceList{}, fmt.Errorf("failed to fetch map data. Status Code: %d", resp.StatusCode)
	}

	if err != nil {
		return PagedResourceList{}, err
	}

	c.cache.Add(reqUrl, body)

	list, err := unmarshalJSON[PagedResourceList](body)

	return list, err
}
