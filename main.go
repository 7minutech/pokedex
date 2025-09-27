package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if moreTokens := scanner.Scan(); !moreTokens {
			log.Fatal(scanner.Err())
		}
		userInput := scanner.Text()
		cleanInput := cleanInput(userInput)
		output := fmt.Sprintf("Your command was: %v", cleanInput[0])
		fmt.Println(output)
	}
}

func cleanInput(text string) []string {
	trimed_text := strings.TrimSpace(text)
	lower_text := strings.ToLower(trimed_text)
	return strings.Split(lower_text, " ")
}
