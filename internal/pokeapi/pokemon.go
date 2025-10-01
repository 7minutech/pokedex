package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
)

type Pokemon struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
	BaseExperience int    `json:"base_experience"`
}

type Stat struct {
	BaseStat int        `json:"base_stat"`
	Stat     StatDetail `json:"stat"`
}

type StatDetail struct {
	Name string `json:"name"`
}

type Type struct {
	Type TypeDetail `json:"type"`
}

type TypeDetail struct {
	Name string `json:"name"`
}

const baseUrl = "https://pokeapi.co/api/v2/pokemon/"

func GetPokemon(name string) (Pokemon, error) {
	fullUrl := baseUrl + name
	var pokemon Pokemon
	if value, found := cache.Get(name); found {
		if err := json.Unmarshal(value, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	resp, err := http.Get(fullUrl)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error: Get request to pokemon -> %w", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error: reading resp.Body -> %w", err)
	}
	cache.Add(fullUrl, data)
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("error: unmarshaling data -> %w", err)
	}
	return pokemon, nil
}

func captureChance(pokemon Pokemon) (chance float64) {
	switch {
	case pokemon.BaseExperience < 25:
		return 0.9
	case pokemon.BaseExperience < 50:
		return 0.8
	case pokemon.BaseExperience < 75:
		return 0.7
	case pokemon.BaseExperience < 100:
		return 0.6
	case pokemon.BaseExperience < 150:
		return 0.5
	case pokemon.BaseExperience < 200:
		return 0.4
	case pokemon.BaseExperience < 300:
		return 0.3
	case pokemon.BaseExperience < 400:
		return 0.2
	default:
		return 0.1
	}

}

func IsCaptured(pokemon Pokemon) bool {
	randNum := (rand.Float64() * 100) + 1
	return randNum < (captureChance(pokemon) * 100)
}

func PokemonInfo(pokemon Pokemon) string {
	info := fmt.Sprintf("Name: %v\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	info += "Stats:\n"
	for _, stat := range pokemon.Stats {
		info += fmt.Sprintf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	info += "Types:\n"
	for _, pokemonType := range pokemon.Types {
		info += fmt.Sprintf("  - %v\n", pokemonType.Type.Name)
	}
	return info
}
