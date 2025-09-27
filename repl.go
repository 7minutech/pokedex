package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		output := fmt.Sprintf("Your command was: %v", words[0])
		fmt.Println(output)
	}
}

func cleanInput(text string) []string {
	trimed_text := strings.TrimSpace(text)
	lower_text := strings.ToLower(trimed_text)
	return strings.Split(lower_text, " ")
}
