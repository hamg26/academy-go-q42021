package repository_test

import (
	"errors"
	"testing"

	"github.com/hamg26/academy-go-q42021/domain/model"
	ir "github.com/hamg26/academy-go-q42021/interface/repository"
	"github.com/hamg26/academy-go-q42021/testutil"
	ur "github.com/hamg26/academy-go-q42021/usecase/repository"
	"github.com/stretchr/testify/assert"
)

func TestPokemonRepository_FindAll(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) ur.PokemonRepository
		assert  func(t *testing.T, p []*model.Pokemon, err error)
	}{
		"success": {
			arrange: func(t *testing.T) ur.PokemonRepository {
				mycsv := new(testutil.MyCsvMock)
				mycsv.On("FindAll").Return(testutil.GetPokemonsRecords(), nil)
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				return r
			},
			assert: func(t *testing.T, p []*model.Pokemon, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 2, len(p))
			},
		},
		"empty": {
			arrange: func(t *testing.T) ur.PokemonRepository {
				mycsv := new(testutil.MyCsvMock)
				records := testutil.GetPokemonsRecords()
				records[1] = nil
				mycsv.On("FindAll").Return(records, nil)
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				return r
			},
			assert: func(t *testing.T, p []*model.Pokemon, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 2, len(p))
			},
		},
		"invalid id from csv": {
			arrange: func(t *testing.T) ur.PokemonRepository {
				mycsv := new(testutil.MyCsvMock)
				records := testutil.GetPokemonsRecords()
				records[0][0] = "asd"
				mycsv.On("FindAll").Return(records, nil)
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				return r
			},
			assert: func(t *testing.T, p []*model.Pokemon, err error) {
				assert.NotEmpty(t, err)
			},
		},
		"error": {
			arrange: func(t *testing.T) ur.PokemonRepository {
				mycsv := new(testutil.MyCsvMock)
				mycsv.On("FindAll").Return(nil, errors.New("fake error"))
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				return r
			},
			assert: func(t *testing.T, p []*model.Pokemon, err error) {
				assert.NotEmpty(t, err)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			r := tt.arrange(t)
			err, p := r.FindAll()
			tt.assert(t, p, err)
		})
	}
}

func TestPokemonRepository_FindOne(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (ur.PokemonRepository, uint64)
		assert  func(t *testing.T, p *model.Pokemon, err error)
	}{
		"success": {
			arrange: func(t *testing.T) (ur.PokemonRepository, uint64) {
				mycsv := new(testutil.MyCsvMock)
				mycsv.On("FindAll").Return(testutil.GetPokemonsRecords(), nil)
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				id := uint64(1)
				return r, id
			},
			assert: func(t *testing.T, p *model.Pokemon, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, p)
			},
		},
		"id not found": {
			arrange: func(t *testing.T) (ur.PokemonRepository, uint64) {
				mycsv := new(testutil.MyCsvMock)
				mycsv.On("FindAll").Return(testutil.GetPokemonsRecords(), nil)
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				id := uint64(0)
				return r, id
			},
			assert: func(t *testing.T, p *model.Pokemon, err error) {
				assert.Nil(t, err)
				assert.Nil(t, p)
			},
		},
		"invalid id from csv": {
			arrange: func(t *testing.T) (ur.PokemonRepository, uint64) {
				mycsv := new(testutil.MyCsvMock)
				record := testutil.GetPokemonsRecords()
				record[0][0] = "asd"
				mycsv.On("FindAll").Return(record, nil)
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				id := uint64(2)
				return r, id
			},
			assert: func(t *testing.T, p *model.Pokemon, err error) {
				assert.NotEmpty(t, err)
				assert.Nil(t, p)
			},
		},
		"error": {
			arrange: func(t *testing.T) (ur.PokemonRepository, uint64) {
				mycsv := new(testutil.MyCsvMock)
				mycsv.On("FindAll").Return(nil, errors.New("fake error"))
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				id := uint64(1)
				return r, id
			},
			assert: func(t *testing.T, p *model.Pokemon, err error) {
				assert.NotEmpty(t, err)
				assert.Nil(t, p)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			r, id := tt.arrange(t)
			err, p := r.FindOne(id)
			tt.assert(t, p, err)
		})
	}
}

func TestPokemonRepository_FindOneDetails(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (ur.PokemonRepository, string)
		assert  func(t *testing.T, p *model.PokemonDetails, err error)
	}{
		"success": {
			arrange: func(t *testing.T) (ur.PokemonRepository, string) {
				mycsv := new(testutil.MyCsvMock)
				api := new(testutil.PokeApiClientMock)
				api.On("GetPokemon", "1").Return(testutil.GetPokemonDetails(), nil)
				r := ir.NewPokemonRepository(mycsv, api)
				id := "1"
				return r, id
			},
			assert: func(t *testing.T, p *model.PokemonDetails, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, p)
			},
		},
		"error": {
			arrange: func(t *testing.T) (ur.PokemonRepository, string) {
				mycsv := new(testutil.MyCsvMock)
				api := new(testutil.PokeApiClientMock)
				api.On("GetPokemon", "1").Return(nil, errors.New("fake API error"))
				r := ir.NewPokemonRepository(mycsv, api)
				id := "1"
				return r, id
			},
			assert: func(t *testing.T, p *model.PokemonDetails, err error) {
				assert.NotEmpty(t, err)
				assert.Nil(t, p)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			r, id := tt.arrange(t)
			err, p := r.FindOneDetails(id)
			tt.assert(t, p, err)
		})
	}
}

func TestPokemonRepository_SavePokemon(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (ur.PokemonRepository, *model.PokemonDetails)
		assert  func(t *testing.T, err error)
	}{
		"success": {
			arrange: func(t *testing.T) (ur.PokemonRepository, *model.PokemonDetails) {
				mycsv := new(testutil.MyCsvMock)
				mycsv.On("Save").Return(nil)
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				return r, testutil.GetPokemonDetails()
			},
			assert: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		"empty_types": {
			arrange: func(t *testing.T) (ur.PokemonRepository, *model.PokemonDetails) {
				mycsv := new(testutil.MyCsvMock)
				mycsv.On("Save").Return(nil)
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				pd := testutil.GetPokemonDetails()
				pd.Types = nil
				return r, pd
			},
			assert: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		"error": {
			arrange: func(t *testing.T) (ur.PokemonRepository, *model.PokemonDetails) {
				mycsv := new(testutil.MyCsvMock)
				mycsv.On("Save").Return(errors.New("fake error"))
				api := new(testutil.PokeApiClientMock)
				r := ir.NewPokemonRepository(mycsv, api)
				return r, testutil.GetPokemonDetails()
			},
			assert: func(t *testing.T, err error) {
				assert.NotEmpty(t, err)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			r, pd := tt.arrange(t)
			err := r.SavePokemon(pd)
			tt.assert(t, err)
		})
	}
}
