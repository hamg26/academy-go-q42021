package testutil

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

// Mocked http client
type HttpClientMock struct {
	mock.Mock
}

// Mocked HttpClientMock.Do
func (ac HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	args := ac.Called(mock.Anything)
	if args.Get(0) != nil {
		return args.Get(0).(*http.Response), args.Error(1)
	}
	return nil, args.Error(1)
}
