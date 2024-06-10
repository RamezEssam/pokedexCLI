package commands

import (
	"fmt"
	"github.com/RamezEssam/pokedexcli/internal/pokecache"
	"github.com/RamezEssam/pokedexcli/internal/api"
)


type MapCommand struct {
	Name string
	Description string
	Callback    func(string, pokecache.Cache) error
}


func CommandMap(offset string, c pokecache.Cache) error {
	
	locations, err := api.GetLocations(offset ,c)
	if err != nil {
		return err
	}

	for _,location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}