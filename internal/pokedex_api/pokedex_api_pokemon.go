package pokedex_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AreaLocation struct {
	PokemonEncounters []Encounter `json:"pokemon_counters"`
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetEcounters(name string) ([]Encounter, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	fullURL := baseURL + name
	resp, err := http.Get(fullURL)
	if err != nil {
		return make([]Encounter, 0), fmt.Errorf("error: Get request to location area -> %w", err)
	}
	var areaLoc AreaLocation
	if err := json.NewDecoder(resp.Body).Decode(&areaLoc); err != nil {
		return make([]Encounter, 0), fmt.Errorf("error: decoding response to pokemon -> %w", err)
	}
	return areaLoc.PokemonEncounters, nil

}

func GetPokemons(encounters []Encounter) []Pokemon {
	pokemons := make([]Pokemon, len(encounters))
	for i, encounter := range encounters {
		pokemons[i] = encounter.Pokemon
	}
	return pokemons
}

func GetPokemonNames(pokemons []Pokemon) []string {
	pokemonNames := make([]string, len(pokemons))
	for i, pokemon := range pokemons {
		pokemonNames[i] = pokemon.Name
	}
	return pokemonNames
}
