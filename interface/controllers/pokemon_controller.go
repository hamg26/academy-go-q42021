package controller

import (
	"net/http"
	"strconv"

	"github.com/hamg26/academy-go-q42021/domain/model"
	"github.com/hamg26/academy-go-q42021/usecase/interactor"
)

type PokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

/*
Returns a new instance of the Pokemon controller
Specific implementation of endpoint requests using an interactor
*/
func NewPokemonController(ps interactor.PokemonInteractor) PokemonController {
	return PokemonController{ps}
}

/*
Response with all the pokemons
*/
func (uc PokemonController) GetPokemons(c Context) error {
	var p []*model.Pokemon

	err, p := uc.pokemonInteractor.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

/*
Response with a specific pokemon
Only reads from the pokemons existing locally
Parameters from the request: Id (int)
*/
func (uc PokemonController) GetPokemon(c Context) error {
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

/*
Response with a specific pokemon details
Request to API and saves the pokemon if it's not already saved
Parameters from the request: Id (int)
*/
func (uc PokemonController) GetPokemonDetails(c Context) error {
	var details *model.PokemonDetails

	rawid := c.Param("id")

	id, e := strconv.ParseUint(rawid, 10, 64)
	if e != nil {
		return c.JSON(http.StatusBadRequest, "Id should be an integer")
	}

	err, details := uc.pokemonInteractor.GetOneDetails(rawid)
	if err != nil {
		return err
	}

	if details == nil {
		return c.JSON(http.StatusNotFound, "Pokemon not found")
	}

	_, p := uc.pokemonInteractor.GetOne(id)
	if p == nil {
		uc.pokemonInteractor.SavePokemon(details)
	}

	return c.JSON(http.StatusOK, details)
}
