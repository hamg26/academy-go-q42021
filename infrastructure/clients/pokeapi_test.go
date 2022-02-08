package pokeapi

import (
	"errors"
	"testing"

	models "github.com/hamg26/academy-go-q42021/domain/model"
	"github.com/hamg26/academy-go-q42021/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPokeApiClient_GetPokemon(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (id string, client pokeApiClient)
		assert  func(t *testing.T, err error, result *models.PokemonDetails)
	}{
		"success": {
			arrange: func(t *testing.T) (string, pokeApiClient) {
				client := new(testutil.HttpClientMock)
				client.On("Do", mock.Anything).Return(testutil.GetPokeApiResponse(200), nil)
				api := NewPokeApiClient("", *client)
				return "1", api
			},
			assert: func(t *testing.T, err error, result *models.PokemonDetails) {
				assert.Nil(t, err)
				assert.NotNil(t, result)
				if assert.NotNil(t, result) {
					assert.Equal(t, "Bulbasaur", result.Name)
					assert.Equal(t, uint64(1), result.Id)
				}
			},
		},
		"error": {
			arrange: func(t *testing.T) (string, pokeApiClient) {
				client := new(testutil.HttpClientMock)
				client.On("Do", mock.Anything).Return(nil, errors.New("fake error"))
				api := NewPokeApiClient("", *client)
				return "1", api
			},
			assert: func(t *testing.T, err error, result *models.PokemonDetails) {
				assert.NotNil(t, err)
				assert.Nil(t, result)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			id, api := tt.arrange(t)
			err, results := api.GetPokemon(id)
			tt.assert(t, err, results)
		})
	}
}
