package cmds

import (
	"fmt"
)

func commandPokedex(c *Config, arg string) error {
	if len(Pokedex) == 0 {
		fmt.Println("You have not captured any pokemon")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range Pokedex {
		fmt.Printf(" - %v\n", name)
	}
	return nil
}
