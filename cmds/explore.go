package cmds

import (
	"fmt"
	"strings"

	"github.com/7minutech/pokedex/internal/pokeapi"
)

func commandExplore(c *Config, arg string) error {
	encounters, err := pokeapi.GetEcounters(arg)
	if err != nil {
		fmt.Printf("couldn't find pokemon at %v\n", arg)
		return nil
	}
	pokemons := pokeapi.GetPokemons(encounters)
	pokemonNames := pokeapi.GetPokemonNames(pokemons)
	fmt.Printf("exploring %v...\n", arg)
	fmt.Println(strings.Join(pokemonNames, "\n"))
	return nil
}
