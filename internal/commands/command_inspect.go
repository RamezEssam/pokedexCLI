package commands

import (
	"fmt"
	"github.com/RamezEssam/pokedexcli/internal/entity"
)

type InspectCommand struct {
	Name string
	Description string
	Callback    func(string, map[string]entity.Pokemon) error
}


func CommadInspect(name string, caught map[string]entity.Pokemon) error {
	pokemon, ok := caught[name]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _,stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _,t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}


	return nil
}