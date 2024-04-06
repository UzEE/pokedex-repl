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

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`

	Abilities              []PokemonAbilities     `json:"abilities"`
	PastAbilities          []PokemonPastAbilities `json:"past_abilities"`
	Forms                  []NamedResource        `json:"forms"`
	GameIndices            []VersionGameIndices   `json:"game_indices"`
	HeldItems              []PokemonHeldItem      `json:"held_items"`
	LocationAreaEncounters string                 `json:"location_area_encounters"`
	Moves                  []Moves                `json:"moves"`
	Sprites                Sprites                `json:"sprites"`
	Cries                  PokemonCries           `json:"cries"`
	Species                NamedResource          `json:"species"`
	Stats                  []Stats                `json:"stats"`
	Types                  []PokemonTypes         `json:"types"`
	PastTypes              []PokemonPastTypes     `json:"past_types"`
}

type PokemonAbilities struct {
	Ability  NamedResource `json:"ability"`
	IsHidden bool          `json:"is_hidden"`
	Slot     int           `json:"slot"`
}

type PokemonPastAbilities struct {
	Abilities  []PokemonAbilities `json:"abilities"`
	Generation NamedResource      `json:"generation"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type PokemonHeldItem struct {
	Item           NamedResource            `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Rarity  int           `json:"rarity"`
	Version NamedResource `json:"version"`
}

type VersionGameIndices struct {
	GameIndex int           `json:"game_index"`
	Version   NamedResource `json:"version"`
}

type PokemonMoveVersion struct {
	LevelLearnedAt  int           `json:"level_learned_at"`
	MoveLearnMethod NamedResource `json:"move_learn_method"`
	VersionGroup    NamedResource `json:"version_group"`
}

type Moves struct {
	Move                NamedResource        `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonTypes struct {
	Slot int           `json:"slot"`
	Type NamedResource `json:"type"`
}

type PokemonPastTypes struct {
	Generation NamedResource  `json:"generation"`
	Types      []PokemonTypes `json:"types"`
}

type Stats struct {
	BaseStat int           `json:"base_stat"`
	Effort   int           `json:"effort"`
	Stat     NamedResource `json:"stat"`
}

type PokemonSprites struct {
	FrontDefault          *string `json:"front_default"`
	FrontShiny            *string `json:"front_shiny"`
	FrontFemale           *string `json:"front_female"`
	FrontShinyFemale      *string `json:"front_shiny_female"`
	FrontGray             *string `json:"front_gray"`
	FrontTransparent      *string `json:"front_transparent"`
	FrontShinyTransparent *string `json:"front_shiny_transparent"`

	BackDefault          *string `json:"back_default"`
	BackShiny            *string `json:"back_shiny"`
	BackFemale           *string `json:"back_female"`
	BackShinyFemale      *string `json:"back_shiny_female"`
	BackGray             *string `json:"back_gray"`
	BackTransparent      *string `json:"back_transparent"`
	BackShinyTransparent *string `json:"back_shiny_transparent"`
}

type Sprites struct {
	PokemonSprites
	Other    SpritesOther    `json:"other"`
	Versions SpritesVersions `json:"versions"`
}

type SpritesVersions struct {
	GenerationI    SpritesGenerationI    `json:"generation-i"`
	GenerationIi   SpritesGenerationII   `json:"generation-ii"`
	GenerationIii  SpritesGenerationIII  `json:"generation-iii"`
	GenerationIv   SpritesGenerationIV   `json:"generation-iv"`
	GenerationV    SpritesGenerationV    `json:"generation-v"`
	GenerationVi   SpritesGenerationVI   `json:"generation-vi"`
	GenerationVii  SpritesGenerationVII  `json:"generation-vii"`
	GenerationViii SpritesGenerationVIII `json:"generation-viii"`
}

type SpritesOther struct {
	DreamWorld      SpritesDreamWorld      `json:"dream_world"`
	Home            SpritesHome            `json:"home"`
	OfficialArtwork SpritesOfficialArtwork `json:"official-artwork"`
	Showdown        SpritesShowdown        `json:"showdown"`
}

type SpritesDreamWorld struct {
	PokemonSprites
}

type SpritesHome struct {
	PokemonSprites
}

type SpritesOfficialArtwork struct {
	PokemonSprites
}

type SpritesShowdown struct {
	PokemonSprites
}

type SpritesRedBlue struct {
	PokemonSprites
}

type SpritesYellow struct {
	PokemonSprites
}

type SpritesGenerationI struct {
	RedBlue SpritesRedBlue `json:"red-blue"`
	Yellow  SpritesYellow  `json:"yellow"`
}

type SpritesCrystal struct {
	PokemonSprites
}

type SpritesGold struct {
	PokemonSprites
}

type SpritesSilver struct {
	PokemonSprites
}

type SpritesGenerationII struct {
	Crystal SpritesCrystal `json:"crystal"`
	Gold    SpritesGold    `json:"gold"`
	Silver  SpritesSilver  `json:"silver"`
}

type SpritesEmerald struct {
	PokemonSprites
}

type SpritesFireRedLeafGreen struct {
	PokemonSprites
}

type SpritesRubySapphire struct {
	PokemonSprites
}

type SpritesGenerationIII struct {
	Emerald          SpritesEmerald          `json:"emerald"`
	FireRedLeafGreen SpritesFireRedLeafGreen `json:"firered-leafgreen"`
	RubySapphire     SpritesRubySapphire     `json:"ruby-sapphire"`
}

type SpritesDiamondPearl struct {
	PokemonSprites
}

type SpritesHeartGoldSoulSilver struct {
	PokemonSprites
}

type SpritesPlatinum struct {
	PokemonSprites
}

type SpritesGenerationIV struct {
	DiamondPearl        SpritesDiamondPearl        `json:"diamond-pearl"`
	HeartGoldSoulSilver SpritesHeartGoldSoulSilver `json:"heartgold-soulsilver"`
	Platinum            SpritesPlatinum            `json:"platinum"`
}

type SpritesAnimated struct {
	PokemonSprites
}

type SpritesBlackWhite struct {
	Animated SpritesAnimated `json:"animated"`
	PokemonSprites
}

type SpritesGenerationV struct {
	BlackWhite SpritesBlackWhite `json:"black-white"`
}

type SpritesOmegaRubyAlphaSapphire struct {
	PokemonSprites
}

type SpritesXY struct {
	PokemonSprites
}

type SpritesGenerationVI struct {
	OmegaRubyAlphaSapphire SpritesOmegaRubyAlphaSapphire `json:"omegaruby-alphasapphire"`
	XY                     SpritesXY                     `json:"x-y"`
}

type SpritesIcons struct {
	PokemonSprites
}

type SpritesUltraSunUltraMoon struct {
	PokemonSprites
}

type SpritesGenerationVII struct {
	Icons             SpritesIcons             `json:"icons"`
	UltraSunUltraMoon SpritesUltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}

type SpritesGenerationVIII struct {
	Icons SpritesIcons `json:"icons"`
}
