package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	PokemonEncounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

func GetEcounters(name string) ([]Encounter, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	fullURL := baseURL + name
	var locArea LocationArea
	if value, found := cache.Get(fullURL); found {
		if err := json.Unmarshal(value, &locArea); err != nil {
			return nil, fmt.Errorf("error: unmarshaling encounters from cache -> %w", err)
		}
		return locArea.PokemonEncounters, nil
	}
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error: Get request to location area -> %w", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error: reading response body -> %w", err)
	}
	cache.Add(fullURL, data)
	if err := json.Unmarshal(data, &locArea); err != nil {
		return nil, fmt.Errorf("error: unmarshaling encounter data -> %w", err)
	}
	return locArea.PokemonEncounters, nil

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
		pokemonNames[i] = "- " + pokemon.Name
	}
	return pokemonNames
}
