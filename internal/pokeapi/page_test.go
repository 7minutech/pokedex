package pokeapi

import (
	"testing"
)

func TestGetLocations(t *testing.T) {
	baseURL := "https://pokeapi.co/api/v2/location-area"
	next := "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"
	cases := []struct {
		input    string
		expected LocationAreaPage
	}{
		{
			input: baseURL,
			expected: LocationAreaPage{
				Next:     &next,
				Previous: nil,
				Results: []location{
					{Name: "canalave-city-area", URL: "https://pokeapi.co/api/v2/location-area/1/"},
					{Name: "eterna-city-area", URL: "https://pokeapi.co/api/v2/location-area/2/"},
					{Name: "pastoria-city-area", URL: "https://pokeapi.co/api/v2/location-area/3/"},
					{Name: "sunyshore-city-area", URL: "https://pokeapi.co/api/v2/location-area/4/"},
					{Name: "sinnoh-pokemon-league-area", URL: "https://pokeapi.co/api/v2/location-area/5/"},
					{Name: "oreburgh-mine-1f", URL: "https://pokeapi.co/api/v2/location-area/6/"},
					{Name: "oreburgh-mine-b1f", URL: "https://pokeapi.co/api/v2/location-area/7/"},
					{Name: "valley-windworks-area", URL: "https://pokeapi.co/api/v2/location-area/8/"},
					{Name: "eterna-forest-area", URL: "https://pokeapi.co/api/v2/location-area/9/"},
					{Name: "fuego-ironworks-area", URL: "https://pokeapi.co/api/v2/location-area/10/"},
					{Name: "mt-coronet-1f-route-207", URL: "https://pokeapi.co/api/v2/location-area/11/"},
					{Name: "mt-coronet-2f", URL: "https://pokeapi.co/api/v2/location-area/12/"},
					{Name: "mt-coronet-3f", URL: "https://pokeapi.co/api/v2/location-area/13/"},
					{Name: "mt-coronet-exterior-snowfall", URL: "https://pokeapi.co/api/v2/location-area/14/"},
					{Name: "mt-coronet-exterior-blizzard", URL: "https://pokeapi.co/api/v2/location-area/15/"},
					{Name: "mt-coronet-4f", URL: "https://pokeapi.co/api/v2/location-area/16/"},
					{Name: "mt-coronet-4f-small-room", URL: "https://pokeapi.co/api/v2/location-area/17/"},
					{Name: "mt-coronet-5f", URL: "https://pokeapi.co/api/v2/location-area/18/"},
					{Name: "mt-coronet-6f", URL: "https://pokeapi.co/api/v2/location-area/19/"},
					{Name: "mt-coronet-1f-from-exterior", URL: "https://pokeapi.co/api/v2/location-area/20/"},
				},
			},
		},
	}

	for _, c := range cases {
		actual, err := GetLocations(c.input)
		if err != nil {
			t.Errorf("GetLocations(%v) = %v; want %v", c.input, err, nil)
		}
		if *actual.Next != *c.expected.Next {
			t.Errorf("*actual.Next = %v; want %v", *actual.Next, *c.expected.Next)
		}
		if actual.Previous != c.expected.Previous {
			t.Errorf("actual.Next = %v; want %v", actual.Next, c.expected.Next)
		}
		if len(actual.Results) != len(c.expected.Results) {
			t.Errorf("actual.Results = %v; want %v", actual.Results, c.expected.Results)
		}
		for i := range actual.Results {
			if actual.Results[i] != c.expected.Results[i] {
				t.Errorf("actual.Results[%v] = %v; want %v", i, actual.Results[i], c.expected.Results[i])
			}
		}
	}
}

func TestLocations(t *testing.T) {
	nextUrl := "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"
	cases := []struct {
		input    LocationAreaPage
		expected []string
	}{
		{
			input: LocationAreaPage{
				Next:     &(nextUrl),
				Previous: nil,
				Results: []location{
					{Name: "canalave-city-area", URL: "https://pokeapi.co/api/v2/location-area/1/"},
					{Name: "eterna-city-area", URL: "https://pokeapi.co/api/v2/location-area/2/"},
					{Name: "pastoria-city-area", URL: "https://pokeapi.co/api/v2/location-area/3/"},
					{Name: "sunyshore-city-area", URL: "https://pokeapi.co/api/v2/location-area/4/"},
					{Name: "sinnoh-pokemon-league-area", URL: "https://pokeapi.co/api/v2/location-area/5/"},
					{Name: "oreburgh-mine-1f", URL: "https://pokeapi.co/api/v2/location-area/6/"},
					{Name: "oreburgh-mine-b1f", URL: "https://pokeapi.co/api/v2/location-area/7/"},
					{Name: "valley-windworks-area", URL: "https://pokeapi.co/api/v2/location-area/8/"},
					{Name: "eterna-forest-area", URL: "https://pokeapi.co/api/v2/location-area/9/"},
					{Name: "fuego-ironworks-area", URL: "https://pokeapi.co/api/v2/location-area/10/"},
					{Name: "mt-coronet-1f-route-207", URL: "https://pokeapi.co/api/v2/location-area/11/"},
					{Name: "mt-coronet-2f", URL: "https://pokeapi.co/api/v2/location-area/12/"},
					{Name: "mt-coronet-3f", URL: "https://pokeapi.co/api/v2/location-area/13/"},
					{Name: "mt-coronet-exterior-snowfall", URL: "https://pokeapi.co/api/v2/location-area/14/"},
					{Name: "mt-coronet-exterior-blizzard", URL: "https://pokeapi.co/api/v2/location-area/15/"},
					{Name: "mt-coronet-4f", URL: "https://pokeapi.co/api/v2/location-area/16/"},
					{Name: "mt-coronet-4f-small-room", URL: "https://pokeapi.co/api/v2/location-area/17/"},
					{Name: "mt-coronet-5f", URL: "https://pokeapi.co/api/v2/location-area/18/"},
					{Name: "mt-coronet-6f", URL: "https://pokeapi.co/api/v2/location-area/19/"},
					{Name: "mt-coronet-1f-from-exterior", URL: "https://pokeapi.co/api/v2/location-area/20/"},
				},
			},
			expected: []string{
				"canalave-city-area",
				"eterna-city-area",
				"pastoria-city-area",
				"sunyshore-city-area",
				"sinnoh-pokemon-league-area",
				"oreburgh-mine-1f",
				"oreburgh-mine-b1f",
				"valley-windworks-area",
				"eterna-forest-area",
				"fuego-ironworks-area",
				"mt-coronet-1f-route-207",
				"mt-coronet-2f",
				"mt-coronet-3f",
				"mt-coronet-exterior-snowfall",
				"mt-coronet-exterior-blizzard",
				"mt-coronet-4f",
				"mt-coronet-4f-small-room",
				"mt-coronet-5f",
				"mt-coronet-6f",
				"mt-coronet-1f-from-exterior",
			},
		},
	}

	for _, c := range cases {
		acutal := Locations(c.input)
		if len(c.input.Results) != len(c.expected) {
			t.Errorf("Locations(%v) = %v; want %v", c.input, acutal, c.expected)
		}
		for i := range acutal {
			if acutal[i] != c.expected[i] {
				t.Errorf("actual[%v] = %v; want %v", i, acutal[i], c.expected[i])
			}
		}
	}
}
