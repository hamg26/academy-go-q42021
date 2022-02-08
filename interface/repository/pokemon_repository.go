package repository

import (
	"log"
	"strconv"

	"github.com/hamg26/academy-go-q42021/domain/model"
	datatstore "github.com/hamg26/academy-go-q42021/infrastructure/datastore"
	"github.com/hamg26/academy-go-q42021/usecase/repository"
)

type pokemonRepository struct {
	mycsv *datatstore.MyCSV
}

func NewPokemonRepository(mycsv *datatstore.MyCSV) repository.PokemonRepository {
	return &pokemonRepository{mycsv}
}

func (pr *pokemonRepository) FindAll() (error, []*model.Pokemon) {
	err, records := pr.mycsv.FindAll()

	if err != nil {
		return err, nil
	}

	var pokemons = make([]*model.Pokemon, len(records))
	for row, content := range records {

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
