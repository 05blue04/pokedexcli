package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	pokecache "github.com/05blue04/pokedexcli/cache"
)

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location on the map",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "view all pokemon you have caught",
			callback: commandPokedex,
		},
	}
}

func main() {
	cfg := Config{
		cache:   pokecache.NewCache(5 * time.Second),
		pokedex: make(map[string]Pokemon),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		input := CleanInput(scanner.Text())

		if len(input) == 0 {
			continue // skip empty lines
		}

		cmd, ok := commands[input[0]]

		if !ok {
			fmt.Println("Unkown command")
			continue
		}

		if err := cmd.callback(&cfg, input[1:]); err != nil {
			fmt.Println("Error: ", err)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
