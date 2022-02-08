package controller

import (
	"net/http"
	"strconv"

	"github.com/hamg26/academy-go-q42021/domain/model"
	"github.com/hamg26/academy-go-q42021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(c Context) error
	GetPokemon(c Context) error
}

func NewPokemonController(ps interactor.PokemonInteractor) PokemonController {
	return &pokemonController{ps}
}

func (uc *pokemonController) GetPokemons(c Context) error {
	var p []*model.Pokemon

	err, p := uc.pokemonInteractor.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func (uc *pokemonController) GetPokemon(c Context) error {
	var p *model.Pokemon

	id, e := strconv.ParseUint(c.Param("id"), 10, 64)
	if e != nil {
		return c.JSON(http.StatusBadRequest, "Id should be an integer")
	}

	err, p := uc.pokemonInteractor.GetOne(id)
	if err != nil {
		return err
	}

	if p == nil {
		return c.JSON(http.StatusNotFound, "Pokemon not found")
	}

	return c.JSON(http.StatusOK, p)
}
