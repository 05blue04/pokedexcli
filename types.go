package main

import (
	"github.com/05blue04/pokedexcli/cache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous *string
	cache pokecache.Cache
}

type ApiResource struct {
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}
