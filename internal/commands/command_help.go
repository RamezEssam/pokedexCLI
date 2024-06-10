package commands

import (
	"fmt"
)

type HelpCommand struct {
	Name string
	Description string
	Callback    func() error
}


func CommandHelp() error{
	fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex REPL
map: Displays the next 20 Pokemon locations
mapb: Displays the previous 20 Pokemon locations
explore <location_area>: Displays all possible pokemons found in the provided <location_area>`)
   return nil
}