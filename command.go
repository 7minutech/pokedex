package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/7minutech/pokedex/internal/pokedex_api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     *string
	previous *string
}

const baseLocationAreaURL = "https://pokeapi.co/api/v2/location-area"

var Commands map[string]cliCommand

var cfg *config

func registerCommands() {
	Commands = map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"help": {name: "help", description: "Displays a help message", callback: commandHelp},
		"map": {name: "map", description: "Displays the names of 20 location areas in the Pokemon world." +
			"\n     Each subsequent call to map displays the next 20 locations",
			callback: commandMap},
	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range Commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *config) error {
	var locArea pokedex_api.LocationArea
	var err error
	if c.next == nil {
		locArea, err = pokedex_api.GetLocations(baseLocationAreaURL)
	} else {
		locArea, err = pokedex_api.GetLocations(*c.next)
	}
	if err != nil {
		log.Fatal("error: getting locations for commandMap", err)
	}
	c.next = locArea.Next
	c.previous = locArea.Previous
	locations := strings.Join(pokedex_api.Locations(locArea), "\n")
	fmt.Println(locations)
	return nil
}
