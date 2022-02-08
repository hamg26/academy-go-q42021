package testutil

import (
	"io"
	"net/http"
	"strings"

	"github.com/hamg26/academy-go-q42021/domain/model"

	"github.com/stretchr/testify/mock"
)

func GetPokeApiResponse(statusCode int) *http.Response {
	body := "{\"height\":7,\"id\":1,\"is_default\":true,\"name\":\"Bulbasaur\",\"order\":1,\"weight\":69,\"base_experience\":64,\"location_area_encounters\":\"https://pokeapi.co/api/v2/pokemon/1/encounters\",\"types\":[{\"slot\":1,\"type\":{\"name\":\"Grass\",\"url\":\"https://pokeapi.co/api/v2/type/12/\"}},{\"slot\":2,\"type\":{\"name\":\"Poison\",\"url\":\"https://pokeapi.co/api/v2/type/4/\"}}],\"abilities\":[{\"ability\":{\"name\":\"overgrow\",\"url\":\"https://pokeapi.co/api/v2/ability/65/\"},\"is_hidden\":false,\"slot\":1},{\"ability\":{\"name\":\"chlorophyll\",\"url\":\"https://pokeapi.co/api/v2/ability/34/\"},\"is_hidden\":true,\"slot\":3}]}"
	return &http.Response{
		Body:       io.NopCloser(strings.NewReader(body)),
		StatusCode: statusCode,
	}
}

type PokeApiClientMock struct {
	mock.Mock
}

func (ac PokeApiClientMock) GetPokemon(id string) (error, *model.PokemonDetails) {
	args := ac.Called(id)
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).(*model.PokemonDetails)
	}
	return args.Error(1), nil
}
