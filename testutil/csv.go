package testutil

import (
	"github.com/stretchr/testify/mock"
)

func GetPokemonsRecords() [][]string {
	return [][]string{
		{"1", "name1", "type1"},
		{"2", "name2", "type2"},
	}
}

type MyCsvMock struct {
	mock.Mock
}

func (mycsv *MyCsvMock) FindAll() (error, [][]string) {
	args := mycsv.Called()
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).([][]string)
	}
	return args.Error(1), nil
}

func (mycsv *MyCsvMock) Save([]string) error {
	args := mycsv.Called()
	return args.Error(0)
}
