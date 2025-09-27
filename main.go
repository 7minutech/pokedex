package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	trimed_text := strings.TrimSpace(text)
	lower_text := strings.ToLower(trimed_text)
	return strings.Split(lower_text, " ")
}
