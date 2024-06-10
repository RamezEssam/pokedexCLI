package api

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/RamezEssam/pokedexcli/internal/pokecache"
)

const (
	LOCATIONS_ENPOINT = "https://pokeapi.co/api/v2/location/"
)

type LocationResponse struct {
    Count    int         `json:"count"`
    Next     string      `json:"next"`
    Previous interface{} `json:"previous"`
    Results  []Location  `json:"results"`
}

type Location struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

func callLocationsAPI(url string) ([]Location, error) {
	res, err := http.Get(url)
	if err != nil {
		return []Location{}, err
	}
	defer res.Body.Close()
	var response LocationResponse
	jsonResp, err := io.ReadAll(res.Body)
	if err != nil {
		return []Location{}, err
	}
	jsonErr := json.Unmarshal(jsonResp, &response)
	if jsonErr != nil {
		return []Location{}, jsonErr
	}

	return response.Results, nil
}


func GetLocations(offset string, c pokecache.Cache) ([]Location, error) {

	url := LOCATIONS_ENPOINT + "?offset=" + offset

	if c.IsEmpty() {
		locations, err := callLocationsAPI(url)
		return locations, err
	}else {
		val, ok := c.Get(url)

		if ok {
			var response LocationResponse
			jsonResp, err := io.ReadAll(val.Body)
			if err != nil {
				return []Location{}, err
			}
			jsonErr := json.Unmarshal(jsonResp, &response)
			if jsonErr != nil {
				return []Location{}, jsonErr
			}

			return response.Results, nil
		}else {
			locations, err := callLocationsAPI(url)
			return locations, err
		}
	}

}