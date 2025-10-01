package cmds

import "fmt"

func commandHelp(c *Config, arg string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range Commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}
