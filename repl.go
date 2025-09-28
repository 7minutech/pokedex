package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var cfg config
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		word := words[0]
		if command, ok := Commands[word]; ok {
			if err := command.callback(&cfg); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("read error:", err)
	}

}

func cleanInput(text string) []string {
	trimed_text := strings.TrimSpace(text)
	lower_text := strings.ToLower(trimed_text)
	return strings.Split(lower_text, " ")
}
