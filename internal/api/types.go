package api

type PagedResourceList struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []NamedResource `json:"results"`
}

type NamedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             NamedResource          `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterVersionDetails struct {
	Rate    int           `json:"rate"`
	Version NamedResource `json:"version"`
}

type EncounterMethodRates struct {
	EncounterMethod NamedResource             `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type Names struct {
	Language NamedResource `json:"language"`
	Name     string        `json:"name"`
}

type Encounter struct {
	MinLevel        int             `json:"min_level"`
	MaxLevel        int             `json:"max_level"`
	ConditionValues []NamedResource `json:"condition_values"`
	Chance          int             `json:"chance"`
	Method          NamedResource   `json:"method"`
}

type VersionEncounterDetails struct {
	EncounterDetails []Encounter   `json:"encounter_details"`
	MaxChance        int           `json:"max_chance"`
	Version          NamedResource `json:"version"`
}

type PokemonEncounters struct {
	Pokemon        NamedResource             `json:"pokemon"`
	VersionDetails []VersionEncounterDetails `json:"version_details"`
}
