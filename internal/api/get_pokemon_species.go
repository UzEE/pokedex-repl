package api

import (
	"fmt"
	"io"

	"github.com/UzEE/pokedexcli/internal/api/types/species"
)

func (c *Client) GetPokemonSpecies(pokemonName string) (species.PokemonSpecies, error) {
	reqUrl := fmt.Sprintf("%s/pokemon-species/%s", baseURL, pokemonName)

	cached, found := c.cache.Get(reqUrl)
	var pokemonSpecies species.PokemonSpecies
	var err error

	if found {
		pokemonSpecies, err = unmarshalJSON[species.PokemonSpecies](cached)
		if err != nil {
			return species.PokemonSpecies{}, fmt.Errorf("failed to cast cache value to Pokémon Species")
		}

		return pokemonSpecies, err
	}

	speciesResp, err := c.client.Get(reqUrl)
	if err != nil {
		return species.PokemonSpecies{}, err
	}

	defer speciesResp.Body.Close()

	speciesBody, err := io.ReadAll(speciesResp.Body)

	if speciesResp.StatusCode >= 400 {
		return species.PokemonSpecies{}, fmt.Errorf("failed to fetch Pokémon Species data. Status Code: %d", speciesResp.StatusCode)
	}

	if err != nil {
		return species.PokemonSpecies{}, err
	}

	c.cache.Add(reqUrl, speciesBody)

	pokemonSpecies, err = unmarshalJSON[species.PokemonSpecies](speciesBody)

	return pokemonSpecies, err
}
