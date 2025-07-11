package main

import (
	"fmt"
	"os"
)
func commandExit(cfg *Config, args []string) error {
	fmt.Fprintf(os.Stdout, "Closing the Pokedex... Goodbye!\n")
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

func commandCatch(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing location area name")
	}

	err := handleCatchCommand(cfg,args[0])

	if err != nil {
		return err
	}

	return nil
}

func commandInspect(cfg *Config, args []string) error {
	poke, ok := cfg.pokedex[args[0]] 
	if !ok{
		return fmt.Errorf("can only inspect pokemons that you have caught")
	}

	fmt.Fprintf(os.Stdout,"Name: %s\nHeight: %d\nWeight: %d\nStats:\n -hp: %d\n -attack: %d\n -defense: %d\n -special-attack: %d\n -special-defense: %d\n -speed: %d\nTypes:\n",
	poke.Name,poke.Height,poke.Weight,poke.Statistics.hp,poke.Statistics.attack,poke.Statistics.defense,poke.Statistics.specialAttack,poke.Statistics.specialDefense,poke.Statistics.speed)
	
	for _, ty := range poke.Types {
		fmt.Fprintf(os.Stdout," - %s\n",ty)
	}
	return nil
}