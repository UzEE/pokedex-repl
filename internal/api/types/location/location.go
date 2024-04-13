package location

import "github.com/UzEE/pokedexcli/internal/api/types"

type Location struct {
	Areas       []types.NamedResource `json:"areas"`
	GameIndices []GameIndices         `json:"game_indices"`
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Names       []Names               `json:"names"`
	Region      Region                `json:"region"`
}

type GameIndices struct {
	GameIndex  int                 `json:"game_index"`
	Generation types.NamedResource `json:"generation"`
}
