package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func processApiData(data []byte, cfg *Config) error {
	var resource ApiResource

	if err := json.Unmarshal(data, &resource); err != nil {
		return err
	}

	cfg.Next = resource.Next
	cfg.Previous = resource.Previous

	for _, location := range resource.Results {
		fmt.Println(location.Name)
	}

	return nil
}
func processLocationData(data []byte) error {

	var resource locationArea

	if err := json.Unmarshal(data, &resource); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemon := range resource.PokemonEncounters {
		fmt.Println("-", pokemon.Pokemon.Name)
	}

	return nil
}

func processPokemonData(data []byte, cfg *Config) error {

	var resource pokemonData

	if err := json.Unmarshal(data, &resource); err != nil {
		return err
	}

	if _, ok := cfg.pokedex[resource.Name]; ok {
		fmt.Printf("You already caught %s!\n", resource.Name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", resource.Name)

	// logic for catching
	exp := float64(resource.BaseExperience)
	catchChance := 0.9 - (exp / 1000)
	if catchChance < 0.1 {
		catchChance = 0.1
	}
	if catchChance > 0.9 {
		catchChance = 0.9
	}

	roll := rand.Float64()
	// fmt.Printf("Catch chance: %.2f | Roll: %.2f\n", catchChance, roll)

	caught := roll < catchChance

	if caught {
		catch := createPokemon(resource)
		fmt.Println(catch)
		cfg.pokedex[resource.Name] = catch
		fmt.Println(resource.Name, "was caught!")
	} else {
		fmt.Println(resource.Name, "escaped!")
	}

	return nil
}
func handleMapCommand(cfg *Config, goBack bool) error {
	url := cfg.Next

	if goBack {
		if cfg.Previous == nil {
			fmt.Println("you're on the first page")
			return nil
		}
		url = *cfg.Previous
	}

	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	if val, check := cfg.cache.Get(url); check {
		return processApiData(val, cfg)
	}

	res, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cfg.cache.Add(url, data)

	return processApiData(data, cfg)
}

func handleExploreCommand(cfg *Config, location string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + location

	if val, check := cfg.cache.Get(url); check {
		return processLocationData(val)
	}
	res, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("error creating request make sure the location name was typed correctly : %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	processLocationData(data)

	cfg.cache.Add(url, data)

	return nil

}

func handleCatchCommand(cfg *Config, pokemon string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	if val, check := cfg.cache.Get(url); check {
		return processPokemonData(val, cfg)
	}

	res, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("error creating request make sure the pokemon name was typed correctly : %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	processPokemonData(data, cfg)

	cfg.cache.Add(url, data)

	return nil
}
