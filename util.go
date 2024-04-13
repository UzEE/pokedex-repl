package main

import (
	"fmt"
	"strings"

	"github.com/UzEE/pokedexcli/internal/api/types/pokemon"
	"github.com/UzEE/pokedexcli/internal/api/types/species"
)

const keyEscape = 27

func decimetersToFeetAndInches(dm int) (int, int) {
	inches := float64(dm) * 3.93701
	feet := int(inches / 12)
	remainder := int(inches) % 12

	return int(feet), int(remainder)
}

func speciesToTerminalColor(c *config, species species.PokemonSpecies) []byte {
	switch species.Color.Name {
	case "black":
		return c.colors.Black
	case "blue":
		return []byte{keyEscape, '[', '9', '4', 'm'}
	case "brown":
		return c.colors.Yellow
	case "gray":
		return []byte{keyEscape, '[', '9', '0', 'm'}
	case "green":
		return c.colors.Green
	case "pink":
		return []byte{keyEscape, '[', '9', '5', 'm'}
	case "purple":
		return c.colors.Magenta
	case "red":
		return c.colors.Red
	case "white":
		return c.colors.White
	case "yellow":
		return []byte{keyEscape, '[', '9', '3', 'm'}
	default:
		return []byte{keyEscape, '[', '9', '7', 'm'}
	}
}

func getFormattedName(c *config, entry pokemon.Pokemon) string {
	return fmt.Sprintf("%s%s%s", speciesToTerminalColor(c, entry.SpeciesEntry), strings.Title(entry.Name), c.colors.Reset)
}
