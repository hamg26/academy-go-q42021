package testutil

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type HttpClientMock struct {
	mock.Mock
}

func (ac HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	args := ac.Called(mock.Anything)
	if args.Get(0) != nil {
		return args.Get(0).(*http.Response), args.Error(1)
	}
	return nil, args.Error(1)
}
