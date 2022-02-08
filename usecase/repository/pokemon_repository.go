package repository

import "github.com/hamg26/academy-go-q42021/domain/model"

type PokemonRepository interface {
	FindAll() (error, []*model.Pokemon)
	FindOne(id uint64) (error, *model.Pokemon)
}
