package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "Charmander",
			expected: []string{"charmander"},
		},
		{
			input:    "",
			expected: []string{""},
		},
		{
			input:    " ",
			expected: []string{""},
		},
		{
			input:    "ChArMander BulbaSAUR PIKACHU squirtle",
			expected: []string{"charmander", "bulbasaur", "pikachu", "squirtle"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%v) = %v; want %v", c.input, actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual[%v] = %v; want %v", i, word, expectedWord)
			}
		}

	}
}
