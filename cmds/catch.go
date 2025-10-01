package cmds

import (
	"fmt"

	"github.com/7minutech/pokedex/internal/pokeapi"
)

func commandCatch(c *Config, arg string) error {
	pokemon, err := pokeapi.GetPokemon(arg)
	if err != nil {
		fmt.Printf("couldn't find pokemon %v\n", arg)
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", arg)
	if pokeapi.IsCaptured(pokemon) {
		fmt.Printf("%v was caught!\n", arg)
		Pokedex[arg] = pokemon
		return nil
	}
	fmt.Printf("%v escaped!\n", arg)
	return nil
}
