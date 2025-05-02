package pokeapi

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Response from the PokeAPI for locations
type LocationAPIResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type LocationPokemonEncountersResponse struct {
	PokemonEncounters []Encounter `json:"pokemon_encounters"`
}
type Encounter struct {
	Pokemon PokemonShallow `json:"pokemon"`
}

type PokemonShallow struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	BaseExperience     int                `json:"base_experience"`
	Height             int                `json:"height"`
	IsDefault          bool               `json:"is_default"`
	Order              int                `json:"order"`
	Weight             int                `json:"weight"`
	Abilities          []PokemonAbility   `json:"abilities"`
	Forms              []NamedAPIResource `json:"forms"`
	GameIndices        []GameIndex        `json:"game_indices"`
	HeldItems          []HeldItem         `json:"held_items"`
	LocationEncounters string             `json:"location_area_encounters"`
	Moves              []PokemonMove      `json:"moves"`
	Species            NamedAPIResource   `json:"species"`
	Sprites            PokemonSprites     `json:"sprites"`
	Cries              PokemonCries       `json:"cries"`
	Stats              []PokemonStat      `json:"stats"`
	Types              []PokemonType      `json:"types"`
	PastTypes          []PastType         `json:"past_types"`
	PastAbilities      []PastAbility      `json:"past_abilities"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type HeldItem struct {
	Item           NamedAPIResource  `json:"item"`
	VersionDetails []HeldItemVersion `json:"version_details"`
}

type HeldItemVersion struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

type PokemonMove struct {
	Move                NamedAPIResource    `json:"move"`
	VersionGroupDetails []MoveVersionDetail `json:"version_group_details"`
}

type MoveVersionDetail struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	Order           int              `json:"order"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type PokemonStat struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PastType struct {
	Generation NamedAPIResource  `json:"generation"`
	Types      []PokemonTypeSlot `json:"types"`
}

type PokemonTypeSlot struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PastAbility struct {
	Generation NamedAPIResource `json:"generation"`
	Abilities  []AbilitySlot    `json:"abilities"`
}

type AbilitySlot struct {
	Ability  *NamedAPIResource `json:"ability"` // nullable
	IsHidden bool              `json:"is_hidden"`
	Slot     int               `json:"slot"`
}

// You can expand this as needed to support all the nested fields
type PokemonSprites struct {
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
	// Add more fields if needed (e.g., shiny, female, versions, etc.)
}
