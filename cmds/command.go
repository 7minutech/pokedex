package cmds

import "github.com/7minutech/pokedex/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	Callback    func(cfg *Config, arg string) error
}

type Config struct {
	next     *string
	previous *string
}

var Commands map[string]cliCommand
var Pokedex map[string]pokeapi.Pokemon

func RegisterCommands() {
	Commands = map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the Pokedex CLI", Callback: commandExit},
		"help": {name: "help", description: "Displays a help message", Callback: commandHelp},
		"map": {name: "map", description: "Displays the names of 20 location areas in the pokemon world." +
			"\n     Each subsequent call to map displays the next 20 locations",
			Callback: commandMap},
		"mapb": {name: "mapb", description: "Displays the names of the last 20 location areas in the Pokemon world." +
			"\n      Must use map at least twice to be able to go back", Callback: commandMapb},
		"explore": {name: "explore <area_name>", description: "Displays a list of all the pokemon located at <area_name>.", Callback: commandExplore},
		"catch":   {name: "catch <pokemon_name>", description: "Attempt to capture <pokemon_name>", Callback: commandCatch},
		"inspect": {name: "inspect <pokemon_name>", description: "Displays the name, height, weight, stats and type(s) of <pokemon_name> if caught", Callback: commandInspect},
		"pokedex": {name: "pokedex", description: "Displays the captured pokemon in pokedex", Callback: commandPokedex},
	}
}

func CreatePokedex() {
	Pokedex = make(map[string]pokeapi.Pokemon)
}
