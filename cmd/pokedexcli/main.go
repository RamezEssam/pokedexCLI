package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/RamezEssam/pokedexcli/internal/commands"
	"github.com/RamezEssam/pokedexcli/internal/entity"
	"github.com/RamezEssam/pokedexcli/internal/pokecache"
)

const START_SCRIPT = "\x1b[32mpokedex > \x1b[0m"


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	offset := 0
	urlCache := pokecache.NewCache(10 * time.Second)
	caught := make(map[string]entity.Pokemon)
	attempted := make(map[string]entity.Pokemon)
	for {
		fmt.Print(START_SCRIPT)
		if !scanner.Scan() {
			break
		}
		input := strings.ToLower(scanner.Text())
		input_fields := strings.Fields(input)
		cmd := input_fields[0]
		switch cmd {
		case "help":
			helpCmd := commands.HelpCommand {
				Name: "help",
				Description: "Displays a help message to the console",
				Callback: commands.CommandHelp,
			}
			err := helpCmd.Callback()
			if err != nil {
				fmt.Println(err)
			}
		case "exit":
			exitCmd := commands.ExitCommand {
				Name: "Exit",
				Description: "Exits the Pokedex REPL",
				Callback: commands.CommandExit,
			}
			err := exitCmd.Callback()
			if err != nil {
				fmt.Println(err)
			}
		case "map":
			mapfCmd := commands.MapCommand {
				Name: "map",
				Description: "Displays the next 20 locations",
				Callback: commands.CommandMap,
			}
			err := mapfCmd.Callback(fmt.Sprint(offset), urlCache)
			if err != nil {
				fmt.Println(err)
			}
			offset += 20
		case "mapb":
			offset -= 20
			if offset < 0 {
				offset = 0
				fmt.Println("No Previous locations")
				continue
			}
			mapbCmd := commands.MapCommand {
				Name: "mapb",
				Description: "Displays the previous 20 locations",
				Callback: commands.CommandMap,
			}
			err := mapbCmd.Callback(fmt.Sprint(offset), urlCache)
			if err != nil {
				fmt.Println(err)
			}
		case "explore":
			if len(input_fields) < 2 {
				fmt.Println("missing location area")
				continue
			}

			location_area := input_fields[1]

			expCmd := commands.ExploreCommand {
				Name: "explore",
				Description: "Displays list of pokemons foudn at given location",
				Callback: commands.CommandExplore,
			}

			err := expCmd.Callback(location_area)
			if err != nil {
				fmt.Println(err)
			}
		case "catch":
			if len(input_fields) < 2 {
				fmt.Println("missing pokemon name")
				continue
			}

			pokemon_name := input_fields[1]

			catchCmd := commands.CatchCommand {
				Name: "catch",
				Description: "Attempts to catch the given pokemon",
				Callback: commands.CommandCatch,
			}

			err := catchCmd.Callback(pokemon_name, caught, attempted)
			if err != nil {
				fmt.Println(err)
			}
		case "inspect":
			if len(input_fields) < 2 {
				fmt.Println("missing pokemon name")
				continue
			}
			pokemon_name := input_fields[1]

			inspctCmd := commands.InspectCommand {
				Name: "inspect",
				Description: "Displays stats for given pokemon",
				Callback: commands.CommadInspect,
			}

			err := inspctCmd.Callback(pokemon_name, caught)
			if err != nil {
				fmt.Println(err)
			}
		case "pokedex":
			poedexCmd := commands.PokedexCommand {
				Name: "pokedex",
				Description: "Displays your caught pokemons",
				Callback: commands.CommandPokedex,
			}

			err := poedexCmd.Callback(caught)
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Printf("Unknown Command: %v\n", input)

		}
	}
	
	
}