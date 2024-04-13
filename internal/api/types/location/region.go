package location

import "github.com/UzEE/pokedexcli/internal/api/types"

type Region struct {
	ID             int                   `json:"id"`
	Locations      []types.NamedResource `json:"locations"`
	MainGeneration types.NamedResource   `json:"main_generation"`
	Name           string                `json:"name"`
	Names          []Names               `json:"names"`
	Pokedexes      []types.NamedResource `json:"pokedexes"`
	VersionGroups  []types.NamedResource `json:"version_groups"`
}
