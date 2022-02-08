package testutil

import (
	"testing"

	controllers "github.com/hamg26/academy-go-q42021/interface/controllers"
)

type context struct {
	FakeError  error
	Params     map[string]string
	StatusCode int
	store      map[string]interface{}
}

func (c *context) JSON(code int, i interface{}) error {
	c.store["Response"] = i
	c.store["StatusCode"] = code
	return c.FakeError
}

func (c *context) Bind(i interface{}) error {
	return c.FakeError
}

func (c *context) Param(p string) string {
	return c.Params[p]
}

func (c *context) Get(key string) interface{} {
	return c.store[key]
}

func (c *context) Set(key string, val interface{}) {
	if c.store == nil {
		c.store = make(map[string]interface{})
	}
	c.store[key] = val
}

func NewContextMock(t *testing.T, fakeError error, paramsValues map[string]string) controllers.Context {
	t.Helper()
	store := map[string]interface{}{
		"Response":   nil,
		"StatusCode": nil,
	}
	return &context{Params: paramsValues, store: store, FakeError: fakeError}
}
