package api

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocationArea(url *string) (PagedResourceList, error) {
	reqUrl := "https://pokeapi.co/api/v2/location-area"

	if url != nil {
		reqUrl = *url
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

	list := PagedResourceList{}
	err = json.Unmarshal(body, &list)

	return list, err
}
