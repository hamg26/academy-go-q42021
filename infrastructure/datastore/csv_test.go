package datastore_test

import (
	"strconv"
	"testing"

	"github.com/hamg26/academy-go-q42021/infrastructure/datastore"
	"github.com/stretchr/testify/assert"
)

func isOdd(n string) bool {
	i, _ := strconv.Atoi(n)
	return i%2 != 0
}

func TestMyCSV_FindAll(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (err error, results [][]string)
		assert  func(t *testing.T, results [][]string, err error)
	}{
		"filter_none_items_all": {
			arrange: func(t *testing.T) (error, [][]string) {
				myCSV := datastore.NewCSV("../../testutil/test_pokemons.csv")
				return myCSV.FindAll("", -1, -1)
			},
			assert: func(t *testing.T, results [][]string, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, results)
				assert.Equal(t, 10, len(results))
			},
		},
		"filter_even_items_one": {
			arrange: func(t *testing.T) (error, [][]string) {
				myCSV := datastore.NewCSV("../../testutil/test_pokemons.csv")
				return myCSV.FindAll("even", 1, 1)
			},
			assert: func(t *testing.T, results [][]string, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, results)
				assert.Equal(t, 1, len(results))
				assert.Equal(t, false, isOdd(results[0][0]))
			},
		},
		"filter_odd_items_30": {
			arrange: func(t *testing.T) (error, [][]string) {
				myCSV := datastore.NewCSV("../../testutil/test_pokemons.csv")
				return myCSV.FindAll("odd", 30, 1)
			},
			assert: func(t *testing.T, results [][]string, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, results)
				assert.Equal(t, 5, len(results))
				assert.Equal(t, true, isOdd(results[0][0]))
			},
		},
		"filter_none_items_30_itemsPerWorker_1": {
			arrange: func(t *testing.T) (error, [][]string) {
				myCSV := datastore.NewCSV("../../testutil/test_pokemons.csv")
				return myCSV.FindAll("", 30, 1)
			},
			assert: func(t *testing.T, results [][]string, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, results)
				assert.Equal(t, 10, len(results))
			},
		},
		"filter_none_items_30_itemsPerWorker_-1": {
			arrange: func(t *testing.T) (error, [][]string) {
				myCSV := datastore.NewCSV("../../testutil/test_pokemons.csv")
				return myCSV.FindAll("", 30, -1)
			},
			assert: func(t *testing.T, results [][]string, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, results)
				assert.Equal(t, 10, len(results))
			},
		},
		"filter_none_items_2_itemsPerWorker_3": {
			arrange: func(t *testing.T) (error, [][]string) {
				myCSV := datastore.NewCSV("../../testutil/test_pokemons.csv")
				return myCSV.FindAll("", 2, 3)
			},
			assert: func(t *testing.T, results [][]string, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, results)
				assert.Equal(t, 2, len(results))
			},
		},
		"filter_none_items_5_itemsPerWorker_2": {
			arrange: func(t *testing.T) (error, [][]string) {
				myCSV := datastore.NewCSV("../../testutil/test_pokemons.csv")
				return myCSV.FindAll("", 5, 2)
			},
			assert: func(t *testing.T, results [][]string, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, results)
				assert.Equal(t, 4, len(results))
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			err, results := tt.arrange(t)
			tt.assert(t, results, err)
		})
	}
}
