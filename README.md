# Pokédex CLI

A simple command-line Pokédex written in Go.
This project uses the [PokéAPI](https://pokeapi.co/)

# Features

- Search for Pokémon by name
- View Pokémon types, stats, and locations
- Make HTTP GET requests to fetch data from the PokéAPI
- JSON parsing for API responses
- Caching to speed up repeated lookups
- Interactive REPL-based command-line interface

# Learning Goals
- Learn how to parse JSON in Go
- Practice making HTTP requests in Go
- Build a CLI tool that makes interacting with a back-end server easier
- Gain hands-on practice with local Go development and tooling
- Understand caching and how to use it to improve performance

# Getting Started
## Prerequisites
- Go 1.25.1 or newer installed (recommended)
- Internet connection (for API requests)

## Clone the Repository
```
git clone https://github.com/yourusername/pokedex-cli.git
cd pokedex-cli
```

## Run the Project
```
go run .
```

## Example Usage
```
Pokedex > help
Welcome to the Pokedex!
Usage:

explore <area_name>: Displays a list of all the pokemon located at <area_name>.
catch <pokemon_name>: Attempt to capture <pokemon_name>
inspect <pokemon_name>: Displays the name, height, weight, stats and type(s) of <pokemon_name> if caught
pokedex: Displays the captured pokemon in pokedex
exit: Exit the Pokedex CLI
help: Displays a help message
map: Displays the names of 20 location areas in the pokemon world.
     Each subsequent call to map displays the next 20 locations
mapb: Displays the names of the last 20 location areas in the Pokemon world.
      Must use map at least twice to be able to go back
```
```
Pokedex > catch pidgey
Throwing a Pokeball at pidgey...
pidgey was caught!
You may now inspect it with the inspect command.
```
```
Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying
```
