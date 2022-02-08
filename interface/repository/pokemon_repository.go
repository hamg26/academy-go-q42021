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
	FindAll(string, int, int) (error, [][]string)
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
	return pr.FindAllConcurrent("", -1, -1)
}

/*
Returns all the Pokemons in the CSV datastore
Uses a workers pool to read and filter the data
*/
func (pr *pokemonRepository) FindAllConcurrent(filter string, items, itemsPerWorker int) (error, []*model.Pokemon) {
	err, records := pr.mycsv.FindAll(filter, items, itemsPerWorker)

	if err != nil {
		return err, nil
	}

	var pokemons = make([]*model.Pokemon, 0)
	for row, content := range records {

		if len(content) == 0 {
			continue
		}

		pokemonId, err := strconv.Atoi(content[0])
		if err != nil {
			log.Println("Unable to parse record", row, err)
		} else {
			p := &model.Pokemon{
				Id:   pokemonId,
				Name: content[1],
				Type: content[2],
			}
			pokemons = append(pokemons, p)
		}
	}

	return nil, pokemons
}

/*
Returns a specific Pokemon from the CSV datastore
*/
func (pr *pokemonRepository) FindOne(id int) (error, *model.Pokemon) {
	err, records := pr.mycsv.FindAll("", -1, -1)

	if err != nil {
		return err, nil
	}

	for row, content := range records {

		pokemonId, err := strconv.Atoi(content[0])
		if err != nil {
			log.Println("Unable to parse record", row, err)
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
func (pr *pokemonRepository) FindOneDetails(id int) (error, *model.PokemonDetails) {
	err, p := pr.api.GetPokemon(strconv.Itoa(id))
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
	record := []string{strconv.Itoa(p.Id), p.Name, t}
	err := pr.mycsv.Save(record)

	return err
}
