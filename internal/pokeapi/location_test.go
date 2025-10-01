package pokeapi

import (
	"testing"
)

func TestGetEcounters(t *testing.T) {
	cases := []struct {
		input    string
		expected []Encounter
	}{
		{
			input: "canalave-city-area",
			expected: []Encounter{
				{Pokemon: PokemonEncounter{Name: "tentacool"}},
				{Pokemon: PokemonEncounter{Name: "tentacruel"}},
				{Pokemon: PokemonEncounter{Name: "staryu"}},
				{Pokemon: PokemonEncounter{Name: "magikarp"}},
				{Pokemon: PokemonEncounter{Name: "gyarados"}},
				{Pokemon: PokemonEncounter{Name: "wingull"}},
				{Pokemon: PokemonEncounter{Name: "pelipper"}},
				{Pokemon: PokemonEncounter{Name: "shellos"}},
				{Pokemon: PokemonEncounter{Name: "gastrodon"}},
				{Pokemon: PokemonEncounter{Name: "finneon"}},
				{Pokemon: PokemonEncounter{Name: "lumineon"}},
			},
		},
	}

	for _, c := range cases {
		actual, err := GetEcounters(c.input)
		if err != nil {
			t.Errorf("GetEncounters(%v) = %v; want %v", actual, err, nil)
		}
		if len(actual) != len(c.expected) {
			t.Errorf("GetEncounters(%v) = %v; want %v", c.input, actual, c.expected)
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("actual[%v]) = %v; want %v", i, actual[i], c.expected[i])
			}
		}
	}
}

func TestGetPokemons(t *testing.T) {
	cases := []struct {
		input    []Encounter
		expected []PokemonEncounter
	}{
		{
			input: []Encounter{
				{Pokemon: PokemonEncounter{Name: "charmander"}},
				{Pokemon: PokemonEncounter{Name: "squirtle"}},
			},
			expected: []PokemonEncounter{{Name: "charmander"}, {Name: "squirtle"}},
		},
	}

	for _, c := range cases {
		actual := GetPokemons(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("GetPokemonNames(%v) = %v; want %v", c.input, actual, c.expected)
		}
		for i := range actual {
			encounter := actual[i]
			expectedEncounter := c.expected[i]
			if encounter != expectedEncounter {
				t.Errorf("actual[%v] = %v; want %v", i, encounter, expectedEncounter)
			}
		}

	}
}

func TestGetPokemonNames(t *testing.T) {
	cases := []struct {
		input    []PokemonEncounter
		expected []string
	}{
		{
			input: []PokemonEncounter{
				{Name: "charmander"},
				{Name: "squirtle"},
			},
			expected: []string{"- charmander", "- squirtle"},
		},
	}

	for _, c := range cases {
		actual := GetPokemonNames(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("GetPokemonNames(%v) = %v; want %v", c.input, actual, c.expected)
		}
		for i := range actual {
			pokemon := actual[i]
			expectedPokemon := c.expected[i]
			if pokemon != expectedPokemon {
				t.Errorf("actual[%v] = %v; want %v", i, pokemon, expectedPokemon)
			}
		}

	}
}
