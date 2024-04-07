package api

import (
	"fmt"
	"io"

	"github.com/UzEE/pokedexcli/internal/api/types"
)

func (c *Client) ListLocationAreas(url *string) (types.PagedResourceList, error) {
	reqUrl := fmt.Sprintf("%s/location-area", baseURL)

	if url != nil {
		reqUrl = *url
	}

	cached, found := c.cache.Get(reqUrl)
	if found {
		list, err := unmarshalJSON[types.PagedResourceList](cached)
		if err != nil {
			return types.PagedResourceList{}, fmt.Errorf("failed to cast cache value to PagedResourceList")
		}

		return list, nil
	}

	resp, err := c.client.Get(reqUrl)
	if err != nil {
		return types.PagedResourceList{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return types.PagedResourceList{}, fmt.Errorf("failed to fetch map data. Status Code: %d", resp.StatusCode)
	}

	if err != nil {
		return types.PagedResourceList{}, err
	}

	c.cache.Add(reqUrl, body)

	list, err := unmarshalJSON[types.PagedResourceList](body)

	return list, err
}
