package model

// Pokemon type item
type PokemonType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// A Pokemon type slot in the array
type PokemonTypeSlot struct {
	Slot int         `json:"slot"`
	Type PokemonType `json:"type"`
}

// Single pokemon ability
type PokemonAbility struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Pokemon abilities information
type PokemonAbilities struct {
	Ability  PokemonAbility `json:"ability"`
	IsHidden bool           `json:"is_hidden"`
	Slot     int            `json:"slot"`
}

// Pokemon details information
type PokemonDetails struct {
	Height                 int                `json:"height"`
	Id                     uint64             `json:"id"`
	IsDefault              bool               `json:"is_default"`
	Name                   string             `json:"name"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	BaseExperience         int                `json:"base_experience"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Types                  []PokemonTypeSlot  `json:"types"`
	Abilities              []PokemonAbilities `json:"abilities"`
}
