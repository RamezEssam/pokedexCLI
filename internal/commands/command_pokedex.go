package commands

import (
	"fmt"
	"github.com/RamezEssam/pokedexcli/internal/entity"
)

type PokedexCommand struct {
	Name string
	Description string
	Callback    func(map[string]entity.Pokemon) error
}

func CommandPokedex(caught map[string]entity.Pokemon) error {
	if len(caught) == 0 {
		fmt.Println("No Pokemons have been caught")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for k := range caught {
		fmt.Printf("  - %s\n", k)
	}
	return nil
}

