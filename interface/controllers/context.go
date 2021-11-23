package controller

type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(string) string
	Get(key string) interface{}
	Set(key string, val interface{})
}
