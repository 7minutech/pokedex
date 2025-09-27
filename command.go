package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var Commands map[string]cliCommand

func registerCommands() {
	Commands = map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"help": {name: "help", description: "Displays a help message", callback: commandHelp},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range Commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}
