package testutil

import (
	"github.com/stretchr/testify/mock"
)

// Returns a list of pokemon records as extracted from a CSV file
func GetPokemonsRecords() [][]string {
	return [][]string{
		{"1", "name1", "type1"},
		{"2", "name2", "type2"},
	}
}

// Mocked MyCsv
type MyCsvMock struct {
	mock.Mock
}

// Mocked MyCsv.FindAll
func (mycsv *MyCsvMock) FindAll(string, int, int) (error, [][]string) {
	args := mycsv.Called()
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).([][]string)
	}
	return args.Error(1), nil
}

// Mocked MyCsv.Save
func (mycsv *MyCsvMock) Save([]string) error {
	args := mycsv.Called()
	return args.Error(0)
}
