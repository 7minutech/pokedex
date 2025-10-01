package cmds

import (
	"fmt"
	"log"
	"strings"

	"github.com/7minutech/pokedex/internal/pokeapi"
)

const baseLocationAreaURL = "https://pokeapi.co/api/v2/location-area"

func commandMap(c *Config, arg string) error {
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

func commandMapb(c *Config, arg string) error {
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
