package species

import "github.com/UzEE/pokedexcli/internal/api/types"

type PokemonSpecies struct {
	ID                   int                       `json:"id"`
	Name                 string                    `json:"name"`
	Order                int                       `json:"order"`
	GenderRate           int                       `json:"gender_rate"`
	CaptureRate          int                       `json:"capture_rate"`
	BaseHappiness        int                       `json:"base_happiness"`
	IsBaby               bool                      `json:"is_baby"`
	IsLegendary          bool                      `json:"is_legendary"`
	IsMythical           bool                      `json:"is_mythical"`
	HatchCounter         int                       `json:"hatch_counter"`
	HasGenderDifferences bool                      `json:"has_gender_differences"`
	FormsSwitchable      bool                      `json:"forms_switchable"`
	GrowthRate           types.NamedResource       `json:"growth_rate"`
	PokedexNumbers       []PokedexNumbers          `json:"pokedex_numbers"`
	EggGroups            []types.NamedResource     `json:"egg_groups"`
	Color                types.NamedResource       `json:"color"`
	Shape                types.NamedResource       `json:"shape"`
	EvolvesFromSpecies   types.NamedResource       `json:"evolves_from_species"`
	EvolutionChain       types.APIResource         `json:"evolution_chain"`
	Habitat              types.NamedResource       `json:"habitat"`
	Generation           types.NamedResource       `json:"generation"`
	Names                []Names                   `json:"names"`
	PalParkEncounters    []PalParkEncounterArea    `json:"pal_park_encounters"`
	FlavorTextEntries    []FlavorTextEntries       `json:"flavor_text_entries"`
	FormDescriptions     []Description             `json:"form_descriptions"`
	Genera               []PokemonGenus            `json:"genera"`
	Varieties            []PokemonSpeciesVarieties `json:"varieties"`
}

type FlavorTextEntries struct {
	FlavorText string              `json:"flavor_text"`
	Language   types.NamedResource `json:"language"`
	Version    types.NamedResource `json:"version"`
}

type PokemonGenus struct {
	Genus    string              `json:"genus"`
	Language types.NamedResource `json:"language"`
}

type Names struct {
	Language types.NamedResource `json:"language"`
	Name     string              `json:"name"`
}

type Description struct {
	Description string              `json:"description"`
	Language    types.NamedResource `json:"language"`
}

type PokedexNumbers struct {
	EntryNumber int                 `json:"entry_number"`
	Pokedex     types.NamedResource `json:"pokedex"`
}

type PokemonSpeciesVarieties struct {
	IsDefault bool                `json:"is_default"`
	Pokemon   types.NamedResource `json:"pokemon"`
}

type PalParkEncounterArea struct {
	BaseScore int                 `json:"base_score"`
	Rate      int                 `json:"rate"`
	Area      types.NamedResource `json:"area"`
}
