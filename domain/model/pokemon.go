package model

type Pokemon struct {
	Id   uint64 `json:"pokemon_id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
