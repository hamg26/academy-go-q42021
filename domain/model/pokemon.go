package model

// Pokemon basic information
type Pokemon struct {
	Id   int    `json:"pokemon_id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
