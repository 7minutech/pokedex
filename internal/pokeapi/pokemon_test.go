package pokeapi

import (
	"testing"
)

func TestGetPokemon(t *testing.T) {
	pokemonName := "pidgey"
	cases := []struct {
		input    string
		expected Pokemon
	}{
		{
			input: "pidgey",
			expected: Pokemon{
				Name:   "pidgey",
				Height: 3,
				Weight: 18,
				Stats: []Stat{
					{BaseStat: 40, Stat: StatDetail{Name: "hp"}},
					{BaseStat: 45, Stat: StatDetail{Name: "attack"}},
					{BaseStat: 40, Stat: StatDetail{Name: "defense"}},
					{BaseStat: 35, Stat: StatDetail{Name: "special-attack"}},
					{BaseStat: 35, Stat: StatDetail{Name: "special-defense"}},
					{BaseStat: 56, Stat: StatDetail{Name: "speed"}},
				},
				Types: []Type{
					{Type: TypeDetail{Name: "normal"}},
					{Type: TypeDetail{Name: "flying"}},
				},
				BaseExperience: 50,
			},
		},
	}

	for _, c := range cases {
		actual, err := GetPokemon(pokemonName)
		if err != nil {
			t.Errorf("GetPokemon(%v) = %v; wanted %v", pokemonName, actual, c.expected)
		}
		if actual.Name != c.expected.Name {
			t.Errorf("GetPokemon(%v) = %v; wanted %v", pokemonName, actual, c.expected.Name)
		}
		if actual.Weight != c.expected.Weight {
			t.Errorf("GetPokemon(%v) = %v; wanted %v", pokemonName, actual, c.expected.Weight)
		}
		if actual.Height != c.expected.Height {
			t.Errorf("GetPokemon(%v) = %v; wanted %v", pokemonName, actual, c.expected.Height)
		}
		if actual.BaseExperience != c.expected.BaseExperience {
			t.Errorf("GetPokemon(%v) = %v; wanted %v", pokemonName, actual, c.expected.BaseExperience)
		}
		if len(actual.Stats) != len(c.expected.Stats) {
			t.Errorf("actual.Stats = %v; wanted %v", actual.Stats, c.expected.Stats)
		}
		for i := range actual.Stats {
			if actual.Stats[i].BaseStat != c.expected.Stats[i].BaseStat ||
				actual.Stats[i].Stat.Name != c.expected.Stats[i].Stat.Name {
				t.Errorf("actual.Stats[%v] = %v; want %v", i, actual.Stats[i], c.expected.Stats[i])
			}
		}
		if len(actual.Types) != len(c.expected.Types) {
			t.Errorf("actual.Types = %v; wanted %v", actual.Types, c.expected.Types)
		}
		for i := range actual.Types {
			if actual.Types[i].Type.Name != c.expected.Types[i].Type.Name {
				t.Errorf("actual.Types[%v] = %v; want %v", i, actual.Types[i], c.expected.Types[i])
			}
		}
	}
}
