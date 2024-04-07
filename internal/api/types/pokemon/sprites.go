package pokemon

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
