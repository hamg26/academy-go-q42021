package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "github.com/hamg26/academy-go-q42021/domain/model"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type pokeApiClient interface {
	GetPokemon(string) (error, *models.PokemonDetails)
}

// Holds the PokeApi internal configuration
type PokeApiClient struct {
	client  httpClient
	BaseUrl string
}

/*
Request to the endpoint /pokemon/{id}
Returns the details of a pokemon
*/
func (pokeApiClient *PokeApiClient) GetPokemon(id string) (err error, result *models.PokemonDetails) {
	err = pokeApiClient.request(fmt.Sprintf("pokemon/%s", id), &result)
	return err, result
}

func (pokeApiClient *PokeApiClient) request(endpoint string, obj interface{}) error {
	req, err := http.NewRequest(http.MethodGet, pokeApiClient.BaseUrl+endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := pokeApiClient.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}

/*
Returns a new instance of the PokeAPI client
*/
func NewPokeApiClient(url string, client httpClient) *PokeApiClient {
	return &PokeApiClient{BaseUrl: url, client: client}
}
