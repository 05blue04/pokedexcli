package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, args []string) error {
	fmt.Fprintf(os.Stdout, "Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args []string) error {
	fmt.Fprintf(os.Stdout, "Welcome to the Pokedex!")
	fmt.Fprintf(os.Stdout, "Usage:\n\n")

	for _, c := range commands {
		fmt.Fprintf(os.Stdout, "%v: %v\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *Config, args []string) error {
	err := handleMapCommand(cfg, false)
	if err != nil {
		return fmt.Errorf("error fetching map data: %w", err)
	}

	return nil
}

func commandMapb(cfg *Config, args []string) error {
	err := handleMapCommand(cfg, true)
	if err != nil {
		return fmt.Errorf("error fetching map data: %w", err)
	}

	return nil
}

func commandExplore(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing location area name")
	}
	
	err := handleExploreCommand(cfg, args[0])

	if err != nil {
		return err
	}

	return nil
}
