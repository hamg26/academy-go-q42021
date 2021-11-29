package repository

import (
	"log"
	"strconv"

	"github.com/hamg26/academy-go-q42021/domain/model"
	"github.com/hamg26/academy-go-q42021/usecase/repository"
)

type pokeApiClient interface {
	GetPokemon(string) (error, *model.PokemonDetails)
}

type myCSV interface {
	FindAll() (error, [][]string)
	Save([]string) error
}

type pokemonRepository struct {
	mycsv myCSV
	api   pokeApiClient
}

/*
Returns a new Pokemon repository
Could use different Datastores/Clients
*/
func NewPokemonRepository(mycsv myCSV, api pokeApiClient) repository.PokemonRepository {
	return &pokemonRepository{mycsv, api}
}

/*
Returns all the Pokemons in the CSV datastore
*/
func (pr *pokemonRepository) FindAll() (error, []*model.Pokemon) {
	err, records := pr.mycsv.FindAll()

	if err != nil {
		return err, nil
	}

	var pokemons = make([]*model.Pokemon, len(records))
	for row, content := range records {

		if len(content) == 0 {
			continue
		}

		pokemonId, err := strconv.ParseUint(content[0], 10, 64)
		if err != nil {
			log.Println("Unable to parse record", row, err)
			return err, nil
		}

		p := &model.Pokemon{
			Id:   pokemonId,
			Name: content[1],
			Type: content[2],
		}
		pokemons[row] = p
	}

	return nil, pokemons
}

/*
Returns a specific Pokemon from the CSV datastore
*/
func (pr *pokemonRepository) FindOne(id uint64) (error, *model.Pokemon) {
	err, records := pr.mycsv.FindAll()

	if err != nil {
		return err, nil
	}

	for row, content := range records {

		pokemonId, err := strconv.ParseUint(content[0], 10, 64)
		if err != nil {
			log.Println("Unable to parse record", row, err)
			return err, nil
		}

		if pokemonId == id {
			p := &model.Pokemon{
				Id:   pokemonId,
				Name: content[1],
				Type: content[2],
			}
			return nil, p
		}
	}

	return nil, nil
}

/*
Returns a specific Pokemon details from the API client
*/
func (pr *pokemonRepository) FindOneDetails(id string) (error, *model.PokemonDetails) {
	err, p := pr.api.GetPokemon(id)
	return err, p
}

/*
Saves a Pokemon to the CSV datastore
*/
func (pr *pokemonRepository) SavePokemon(p *model.PokemonDetails) error {
	var t = ""
	if p.Types != nil || len(p.Types) > 0 {
		t = p.Types[0].Type.Name
	}
	record := []string{strconv.FormatUint(p.Id, 10), p.Name, t}
	err := pr.mycsv.Save(record)

	return err
}
