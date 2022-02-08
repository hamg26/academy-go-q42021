package router

import (
	"log"

	"github.com/hamg26/academy-go-q42021/config"
	"github.com/hamg26/academy-go-q42021/infrastructure/validators"
	controller "github.com/hamg26/academy-go-q42021/interface/controllers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*
Returns a new instance of the Router
Initializes the endpoints
*/
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	config.ReadConfig()

	if config.C.Logging == true {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           "${time_custom} Request method=${method}, uri=${uri}, status=${status}\n",
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           log.Writer(),
		}))
	}
	e.Use(middleware.Recover())

	e.Validator = validators.NewConcurrentParamsValidator(validator.New())
	e.GET("/pokemons", func(context echo.Context) error { return c.GetPokemons(context) })
	e.GET("/pokemons/:id", func(context echo.Context) error { return c.GetPokemon(context) })
	e.GET("/pokemons/:id/details", func(context echo.Context) error { return c.GetPokemonDetails(context) })
	e.GET("/pokemons/concurrent", func(context echo.Context) error { return c.GetPokemonsConcurrent(context) })
	return e
}
