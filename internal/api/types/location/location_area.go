package location

import "github.com/UzEE/pokedexcli/internal/api/types"

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             types.NamedResource    `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethodRates struct {
	EncounterMethod types.NamedResource       `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int                 `json:"rate"`
	Version types.NamedResource `json:"version"`
}

type Names struct {
	Language types.NamedResource `json:"language"`
	Name     string              `json:"name"`
}

type Encounter struct {
	MinLevel        int                   `json:"min_level"`
	MaxLevel        int                   `json:"max_level"`
	ConditionValues []types.NamedResource `json:"condition_values"`
	Chance          int                   `json:"chance"`
	Method          types.NamedResource   `json:"method"`
}

type VersionEncounterDetails struct {
	EncounterDetails []Encounter         `json:"encounter_details"`
	MaxChance        int                 `json:"max_chance"`
	Version          types.NamedResource `json:"version"`
}

type PokemonEncounters struct {
	Pokemon        types.NamedResource       `json:"pokemon"`
	VersionDetails []VersionEncounterDetails `json:"version_details"`
}
