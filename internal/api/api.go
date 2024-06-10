package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/RamezEssam/pokedexcli/internal/entity"
	"github.com/RamezEssam/pokedexcli/internal/pokecache"
)

const (
	LOCATIONS_ENDPOINT = "https://pokeapi.co/api/v2/location/"
	LOCATIONS_AREA_ENDPOINT = "https://pokeapi.co/api/v2/location-area/"
)


func callLocationsAPI(url string) ([]entity.Location, error) {
	res, err := http.Get(url)
	if err != nil {
		return []entity.Location{}, err
	}
	defer res.Body.Close()
	var response entity.LocationResponse
	jsonResp, err := io.ReadAll(res.Body)
	if err != nil {
		return []entity.Location{}, err
	}
	jsonErr := json.Unmarshal(jsonResp, &response)
	if jsonErr != nil {
		return []entity.Location{}, jsonErr
	}

	return response.Results, nil
}


func GetLocations(offset string, c pokecache.Cache) ([]entity.Location, error) {

	url := LOCATIONS_ENDPOINT + "?offset=" + offset

	if c.IsEmpty() {
		locations, err := callLocationsAPI(url)
		c.Add(url, locations)
		return locations, err
	}else {
		val, ok := c.Get(url)

		if ok {
			return val, nil
		}else {
			locations, err := callLocationsAPI(url)
			c.Add(url, locations)
			return locations, err
		}
	}

}


func callLocationsAreaAPI(url string) ([]entity.PokemonEncounter, error) {
	res, err := http.Get(url)
	if err != nil {
		return []entity.PokemonEncounter{}, err
	}
	defer res.Body.Close()
	var response entity.LocationArea
	if res.StatusCode != 200 {
		return []entity.PokemonEncounter{}, errors.New("location area not found")
	}
	jsonResp, err := io.ReadAll(res.Body)
	if err != nil {
		return []entity.PokemonEncounter{}, err
	}
	jsonErr := json.Unmarshal(jsonResp, &response)
	if jsonErr != nil {
		return []entity.PokemonEncounter{}, jsonErr
	}
	return response.PokemonEncounters, nil
}

func GetPokemons(location_area string) ([]entity.PokemonEncounter, error) {
	url := LOCATIONS_AREA_ENDPOINT  + location_area

	pokemons, err := callLocationsAreaAPI(url)

	if err != nil {
		return []entity.PokemonEncounter{}, err
	}

	return pokemons, nil
}