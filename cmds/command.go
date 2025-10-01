package cmds

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

func RegisterCommands() {
	Commands = map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the Pokedex", Callback: commandExit},
		"help": {name: "help", description: "Displays a help message", Callback: commandHelp},
		"map": {name: "map", description: "Displays the names of 20 location areas in the Pokemon world." +
			"\n     Each subsequent call to map displays the next 20 locations",
			Callback: commandMap},
		"mapb": {name: "mapb", description: "Displays the names of the last 20 location areas in the Pokemon world." +
			"\n      Must use map at least twice to be able to go back", Callback: commandMapb},
		"explore": {name: "explore <area_name>", description: "Displays a list of all the Pokémon located at <area_name>.", Callback: commandExplore},
	}
}
