package commands

import (
	"fmt"
	"github.com/RamezEssam/pokedexcli/internal/api"
)

type ExploreCommand struct {
	Name string
	Description string
	Callback    func(string) error
}

func CommandExplore(location_area string) error {
	pokemons, err := api.GetPokemons(location_area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location_area)
	if len(pokemons) > 0 {
		fmt.Println("Found Pokemon:")
	}
	for _,pokemon := range pokemons {
		fmt.Printf("  -%s\n", pokemon.Pokemon.Name)
	}
	return nil
}