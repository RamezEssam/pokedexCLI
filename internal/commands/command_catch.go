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


func CommandCatch(name string, caught map[string]entity.Pokemon, attempted map[string]entity.Pokemon) error {

	attempted_pokemon, ok := attempted[name]

	if ok {
		_, ok := caught[name]
		if ok {
			fmt.Printf("%s already captured!\n", name)
			return nil
		}
		base_exp := attempted_pokemon.BaseExperience

		rand_int := rand.Intn(base_exp)

		probability_caught := 1.0 / float64(rand_int)

		rand_float := rand.Float64()

		pokemon_caught :=  rand_float < probability_caught

		if pokemon_caught {
			fmt.Printf("Throwing a Pokeball at %s...\n", name)
			fmt.Printf("%s was caught!\n", name)
			caught[name] = attempted_pokemon
			return nil
		}else {
			fmt.Printf("Throwing a Pokeball at %s...\n", name)
			fmt.Printf("%s escaped!\n", name)
			return nil
		}
	}else{
		caught_pokemon, err := api.GetPokemonStats(name)

		attempted[name] = caught_pokemon

		if err != nil {
			return err
		}

		base_exp := caught_pokemon.BaseExperience

		rand_int := rand.Intn(base_exp)

		probability_caught := 1.0 / float64(rand_int)

		rand_float := rand.Float64()

		pokemon_caught :=  rand_float < probability_caught

		if pokemon_caught {
			fmt.Printf("Throwing a Pokeball at %s...\n", name)
			fmt.Printf("%s was caught!\n", name)
			caught[name] = caught_pokemon
			return nil
		}else {
			fmt.Printf("Throwing a Pokeball at %s...\n", name)
			fmt.Printf("%s escaped!\n", name)
			return nil
		}

	}
}