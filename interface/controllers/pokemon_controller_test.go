package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/hamg26/academy-go-q42021/domain/model"
	pc "github.com/hamg26/academy-go-q42021/interface/controllers"
	"github.com/hamg26/academy-go-q42021/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPokemonController_GetPokemons(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (pc.PokemonController, pc.Context)
		assert  func(t *testing.T, context pc.Context, err error)
	}{
		"success": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetAll").Return(testutil.GetPokemons(), nil)
				controller := pc.NewPokemonController(pi)
				context := testutil.NewContextMock(t, nil, nil)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusOK, context.Get("StatusCode"))
			},
		},
		"error": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetAll").Return(nil, errors.New("fake error"))
				controller := pc.NewPokemonController(pi)
				context := testutil.NewContextMock(t, nil, nil)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.NotNil(t, err)
				assert.NotNil(t, context)
				assert.Nil(t, context.Get("StatusCode"))
				assert.Nil(t, context.Get("Response"))
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			controller, context := tt.arrange(t)
			err := controller.GetPokemons(context)
			tt.assert(t, context, err)
		})
	}
}

func TestPokemonController_GetPokemon(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (pc.PokemonController, pc.Context)
		assert  func(t *testing.T, context pc.Context, err error)
	}{
		"success": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetOne", uint64(1)).Return(testutil.GetPokemons()[0], nil)
				controller := pc.NewPokemonController(pi)
				params := map[string]string{
					"id": "1",
				}
				context := testutil.NewContextMock(t, nil, params)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusOK, context.Get("StatusCode"))
				response := context.Get("Response")
				if assert.NotNil(t, response) {
					pokemon := response.(*model.Pokemon)
					assert.Equal(t, "name1", pokemon.Name)
					assert.Equal(t, uint64(1), pokemon.Id)
				}
			},
		},
		"not_found": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetOne", uint64(0)).Return(nil, nil)
				controller := pc.NewPokemonController(pi)
				params := map[string]string{
					"id": "0",
				}
				context := testutil.NewContextMock(t, nil, params)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusNotFound, context.Get("StatusCode"))
				response := context.Get("Response")
				if assert.NotNil(t, response) {
					assert.Equal(t, "Pokemon not found", response)
				}
			},
		},
		"invalid_id": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				controller := pc.NewPokemonController(pi)
				params := map[string]string{
					"id": "asd",
				}
				context := testutil.NewContextMock(t, nil, params)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusBadRequest, context.Get("StatusCode"))
				response := context.Get("Response")
				if assert.NotNil(t, response) {
					assert.Equal(t, "Id should be an integer", response)
				}
			},
		},
		"error": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetOne", uint64(1)).Return(testutil.GetPokemons()[0], errors.New("fake error"))
				controller := pc.NewPokemonController(pi)
				params := map[string]string{
					"id": "1",
				}
				context := testutil.NewContextMock(t, nil, params)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.NotNil(t, err)
				assert.NotNil(t, context)
				assert.Nil(t, context.Get("StatusCode"))
				assert.Nil(t, context.Get("Response"))
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			controller, context := tt.arrange(t)
			err := controller.GetPokemon(context)
			tt.assert(t, context, err)
		})
	}
}

func TestPokemonController_GetPokemonDetails(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (pc.PokemonController, pc.Context)
		assert  func(t *testing.T, context pc.Context, err error)
	}{
		"success": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetOne", uint64(1)).Return(testutil.GetPokemons()[0], nil)
				pi.On("GetOneDetails", "1").Return(testutil.GetPokemonDetails(), nil)
				controller := pc.NewPokemonController(pi)
				params := map[string]string{
					"id": "1",
				}
				context := testutil.NewContextMock(t, nil, params)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusOK, context.Get("StatusCode"))
				response := context.Get("Response")
				if assert.NotNil(t, response) {
					pokemon := response.(*model.PokemonDetails)
					assert.Equal(t, "name1", pokemon.Name)
					assert.Equal(t, uint64(1), pokemon.Id)
				}
			},
		},
		"success_saving": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetOne", uint64(1)).Return(nil, nil)
				pi.On("GetOneDetails", "1").Return(testutil.GetPokemonDetails(), nil)
				pi.On("SavePokemon", mock.Anything).Return(nil)
				controller := pc.NewPokemonController(pi)
				params := map[string]string{
					"id": "1",
				}
				context := testutil.NewContextMock(t, nil, params)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusOK, context.Get("StatusCode"))
				response := context.Get("Response")
				if assert.NotNil(t, response) {
					pokemon := response.(*model.PokemonDetails)
					assert.Equal(t, "name1", pokemon.Name)
					assert.Equal(t, uint64(1), pokemon.Id)
				}
			},
		},
		"not_found": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetOneDetails", "0").Return(nil, nil)
				controller := pc.NewPokemonController(pi)
				params := map[string]string{
					"id": "0",
				}
				context := testutil.NewContextMock(t, nil, params)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusNotFound, context.Get("StatusCode"))
				response := context.Get("Response")
				if assert.NotNil(t, response) {
					assert.Equal(t, "Pokemon not found", response)
				}
			},
		},
		"invalid_id": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetOneDetails", "asd").Return(nil, nil)
				controller := pc.NewPokemonController(pi)
				params := map[string]string{
					"id": "asd",
				}
				context := testutil.NewContextMock(t, nil, params)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusBadRequest, context.Get("StatusCode"))
				response := context.Get("Response")
				if assert.NotNil(t, response) {
					assert.Equal(t, "Id should be an integer", response)
				}
			},
		},
		"error": {
			arrange: func(t *testing.T) (pc.PokemonController, pc.Context) {
				pi := new(testutil.PokemonInteractor)
				pi.On("GetOneDetails", "1").Return(nil, errors.New("error saving"))
				pi.On("GetOne", uint64(1)).Return(nil, nil)
				pi.On("SavePokemon", mock.Anything).Return(errors.New("error saving"))
				controller := pc.NewPokemonController(pi)
				params := map[string]string{
					"id": "1",
				}
				context := testutil.NewContextMock(t, nil, params)
				return controller, context
			},
			assert: func(t *testing.T, context pc.Context, err error) {
				assert.NotNil(t, err)
				assert.NotNil(t, context)
				assert.Nil(t, context.Get("StatusCode"))
				assert.Nil(t, context.Get("Response"))
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			controller, context := tt.arrange(t)
			err := controller.GetPokemonDetails(context)
			tt.assert(t, context, err)
		})
	}
}
