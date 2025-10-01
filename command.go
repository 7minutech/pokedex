package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/7minutech/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, arg string) error
}

type config struct {
	next     *string
	previous *string
}

const baseLocationAreaURL = "https://pokeapi.co/api/v2/location-area"

var Commands map[string]cliCommand

func registerCommands() {
	Commands = map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"help": {name: "help", description: "Displays a help message", callback: commandHelp},
		"map": {name: "map", description: "Displays the names of 20 location areas in the Pokemon world." +
			"\n     Each subsequent call to map displays the next 20 locations",
			callback: commandMap},
		"mapb": {name: "mapb", description: "Displays the names of the last 20 location areas in the Pokemon world." +
			"\n      Must use map at least twice to be able to go back", callback: commandMapb},
		"explore": {name: "explore <area_name>", description: "Displays a list of all the Pokémon located at <area_name>.", callback: commandExplore},
	}
}

func commandExit(c *config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, arg string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range Commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *config, arg string) error {
	var locAreaPage pokeapi.LocationAreaPage
	var err error
	if c.next == nil {
		locAreaPage, err = pokeapi.GetLocations(baseLocationAreaURL)
	} else {
		locAreaPage, err = pokeapi.GetLocations(*c.next)
	}
	if err != nil {
		log.Fatal("error: getting locations for commandMap", err)
	}
	c.next = locAreaPage.Next
	c.previous = locAreaPage.Previous
	locations := strings.Join(pokeapi.Locations(locAreaPage), "\n")
	fmt.Println(locations)
	return nil
}

func commandMapb(c *config, arg string) error {
	var locAreaPage pokeapi.LocationAreaPage
	var err error
	if c.next == nil && c.previous == nil {
		fmt.Println("no pages have been listed")
		return nil
	}
	if c.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		locAreaPage, err = pokeapi.GetLocations(*c.previous)
	}
	if err != nil {
		log.Fatal("error: getting locations for commandMap", err)
	}
	c.next = locAreaPage.Next
	c.previous = locAreaPage.Previous
	locations := strings.Join(pokeapi.Locations(locAreaPage), "\n")
	fmt.Println(locations)
	return nil
}

func commandExplore(c *config, arg string) error {
	encounters, err := pokeapi.GetEcounters(arg)
	if err != nil {
		fmt.Printf("couldn't find pokemon at %v\n", arg)
		return nil
	}
	pokemons := pokeapi.GetPokemons(encounters)
	pokemonNames := pokeapi.GetPokemonNames(pokemons)
	fmt.Printf("exploring %v...\n", arg)
	fmt.Println(strings.Join(pokemonNames, "\n"))
	return nil
}
