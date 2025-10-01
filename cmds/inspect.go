package cmds

import (
	"fmt"

	"github.com/7minutech/pokedex/internal/pokeapi"
)

func commandInspect(c *Config, arg string) error {
	pokemon, ok := Pokedex[arg]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Print(pokeapi.PokemonInfo(pokemon))
	return nil
}
