package controller

import (
	"net/http"
	"strconv"

	"github.com/hamg26/academy-go-q42021/domain/model"
	validators "github.com/hamg26/academy-go-q42021/infrastructure/validators"
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
func (uc PokemonController) GetPokemon(c Context) (err error) {
	var p *model.Pokemon

	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, "Id should be an integer")
	}

	err, p = uc.pokemonInteractor.GetOne(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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
func (uc PokemonController) GetPokemonDetails(c Context) (err error) {
	var details *model.PokemonDetails

	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, "Id should be an integer")
	}

	err, details = uc.pokemonInteractor.GetOneDetails(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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

/*
Response with all the pokemons (can be filtered)
Only reads from the pokemons existing locally.
Uses multiple workers depending of the # of items and # of itemsPerWorker
Parameters from the request: type (string), items (int), itemsPerWorker (int)
Where:
 - Type: Is used to filter the IDs, filters available: ["even", "odd", ""]
 - Items: Limit of the items to be returned
 - ItemsPerWorker: Limit of the items to be processed by one worker
*/
func (uc PokemonController) GetPokemonsConcurrent(c Context) (err error) {
	var p []*model.Pokemon

	concurrentParams := new(validators.ConcurrentParams)

	if err = c.Bind(concurrentParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(concurrentParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err, p = uc.pokemonInteractor.GetAllConcurrent(concurrentParams.Type, concurrentParams.Items, concurrentParams.ItemsPerWorker)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
