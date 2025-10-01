package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/7minutech/pokedex/cmds"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var cfg cmds.Config
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		arg := ""
		if len(words) > 1 {
			arg = words[1]
		}
		word := words[0]
		if command, ok := cmds.Commands[word]; ok {
			if err := command.Callback(&cfg, arg); err != nil {
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
