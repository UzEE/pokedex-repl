package api

import (
	"fmt"
	"io"

	"github.com/UzEE/pokedexcli/internal/api/types/pokemon"
)

func (c *Client) GetPokemon(pokemonName string) (pokemon.Pokemon, error) {
	reqUrl := fmt.Sprintf("%s/pokemon/%s", baseURL, pokemonName)

	cached, found := c.cache.Get(reqUrl)
	if found {
		mon, err := unmarshalJSON[pokemon.Pokemon](cached)
		if err != nil {
			return pokemon.Pokemon{}, fmt.Errorf("failed to cast cache value to Pokémon")
		}

		return mon, nil
	}

	resp, err := c.client.Get(reqUrl)
	if err != nil {
		return pokemon.Pokemon{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return pokemon.Pokemon{}, fmt.Errorf("failed to fetch Pokémon data. Status Code: %d", resp.StatusCode)
	}

	if err != nil {
		return pokemon.Pokemon{}, err
	}

	c.cache.Add(reqUrl, body)

	pokemon, err := unmarshalJSON[pokemon.Pokemon](body)

	return pokemon, err
}
