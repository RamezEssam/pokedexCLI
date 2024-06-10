package api

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/RamezEssam/pokedexcli/internal/pokecache"
	"github.com/RamezEssam/pokedexcli/internal/entity"
)

const (
	LOCATIONS_ENPOINT = "https://pokeapi.co/api/v2/location/"
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

	url := LOCATIONS_ENPOINT + "?offset=" + offset

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