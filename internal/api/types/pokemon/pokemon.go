package pokemon

import (
	"github.com/UzEE/pokedexcli/internal/api/types"
	"github.com/UzEE/pokedexcli/internal/api/types/species"
)

type Pokemon struct {
	ID                     int                    `json:"id"`
	Name                   string                 `json:"name"`
	BaseExperience         int                    `json:"base_experience"`
	Height                 int                    `json:"height"`
	IsDefault              bool                   `json:"is_default"`
	Order                  int                    `json:"order"`
	Weight                 int                    `json:"weight"`
	Abilities              []PokemonAbilities     `json:"abilities"`
	PastAbilities          []PokemonPastAbilities `json:"past_abilities"`
	Forms                  []types.NamedResource  `json:"forms"`
	GameIndices            []VersionGameIndices   `json:"game_indices"`
	HeldItems              []PokemonHeldItem      `json:"held_items"`
	LocationAreaEncounters string                 `json:"location_area_encounters"`
	Moves                  []Moves                `json:"moves"`
	Sprites                Sprites                `json:"sprites"`
	Cries                  PokemonCries           `json:"cries"`
	Species                types.NamedResource    `json:"species"`
	Stats                  []Stats                `json:"stats"`
	Types                  []PokemonTypes         `json:"types"`
	PastTypes              []PokemonPastTypes     `json:"past_types"`

	// Additional fields
	SpeciesEntry species.PokemonSpecies
}

type PokemonAbilities struct {
	Ability  types.NamedResource `json:"ability"`
	IsHidden bool                `json:"is_hidden"`
	Slot     int                 `json:"slot"`
}

type PokemonPastAbilities struct {
	Abilities  []PokemonAbilities  `json:"abilities"`
	Generation types.NamedResource `json:"generation"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Stats struct {
	BaseStat int                 `json:"base_stat"`
	Effort   int                 `json:"effort"`
	Stat     types.NamedResource `json:"stat"`
}

type PokemonHeldItem struct {
	Item           types.NamedResource      `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Rarity  int                 `json:"rarity"`
	Version types.NamedResource `json:"version"`
}

type VersionGameIndices struct {
	GameIndex int                 `json:"game_index"`
	Version   types.NamedResource `json:"version"`
}

type PokemonMoveVersion struct {
	LevelLearnedAt  int                 `json:"level_learned_at"`
	MoveLearnMethod types.NamedResource `json:"move_learn_method"`
	VersionGroup    types.NamedResource `json:"version_group"`
}

type Moves struct {
	Move                types.NamedResource  `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonTypes struct {
	Slot int                 `json:"slot"`
	Type types.NamedResource `json:"type"`
}

type PokemonPastTypes struct {
	Generation types.NamedResource `json:"generation"`
	Types      []PokemonTypes      `json:"types"`
}
