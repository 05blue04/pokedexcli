package main

import (
	"time"

	pokecache "github.com/05blue04/pokedexcli/cache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

type Config struct {
	Next     string
	Previous *string
	cache    pokecache.Cache
	pokedex  map[string]Pokemon
}

type Pokemon struct {
	Name           string
	BaseExperience int
	Weight         int
	Height         int
	Types           []string
	Statistics     Stats
	CaughtAt       time.Time
}

type Stats struct {
	hp             int
	attack         int
	defense        int
	specialAttack  int
	specialDefense int
	speed          int
}

type ApiResource struct {
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type pokemonData struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
