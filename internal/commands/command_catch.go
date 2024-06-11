package commands

import (
	"fmt"
	"math/rand"
	"github.com/RamezEssam/pokedexcli/internal/api"
	"github.com/RamezEssam/pokedexcli/internal/entity"
)

type CatchCommand struct {
	Name string
	Description string
	Callback    func(string, map[string]entity.Pokemon, map[string]entity.Pokemon) error
}


func playCatch(attempted_pokemon entity.Pokemon, caught map[string]entity.Pokemon) {

		rand_int := rand.Intn(attempted_pokemon.BaseExperience)

		probability_caught := 1.0 / float64(rand_int)

		rand_float := rand.Float64()

		pokemon_caught :=  rand_float < probability_caught

		if pokemon_caught {
			fmt.Printf("Throwing a Pokeball at %s...\n", attempted_pokemon.Name)
			fmt.Printf("%s was caught!\n", attempted_pokemon.Name)
			caught[attempted_pokemon.Name] = attempted_pokemon
		}else {
			fmt.Printf("Throwing a Pokeball at %s...\n", attempted_pokemon.Name)
			fmt.Printf("%s escaped!\n", attempted_pokemon.Name)
		}
}

func CommandCatch(name string, caught map[string]entity.Pokemon, attempted map[string]entity.Pokemon) error {

	attempted_pokemon, ok := attempted[name]

	if ok {
		_, ok := caught[name]
		if ok {
			fmt.Printf("%s already captured!\n", name)
			return nil
		}
		playCatch(attempted_pokemon, caught)
		return nil
	}else{
		caught_pokemon, err := api.GetPokemonStats(name)

		attempted[name] = caught_pokemon

		if err != nil {
			return err
		}

		playCatch(caught_pokemon, caught)
		return nil

	}
}