package api

import (
	"fmt"
	"io"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	reqUrl := fmt.Sprintf("%s/pokemon/%s", baseURL, pokemonName)

	cached, found := c.cache.Get(reqUrl)
	if found {
		pokemon, err := unmarshalJSON[Pokemon](cached)
		if err != nil {
			return Pokemon{}, fmt.Errorf("failed to cast cache value to Pokémon")
		}

		return pokemon, nil
	}

	resp, err := c.client.Get(reqUrl)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return Pokemon{}, fmt.Errorf("failed to fetch Pokémon data. Status Code: %d", resp.StatusCode)
	}

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(reqUrl, body)

	pokemon, err := unmarshalJSON[Pokemon](body)

	return pokemon, err
}
