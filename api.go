package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	if val, check := cfg.cache.Get(url) ; check{
		return processApiData(val,cfg)
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

	return processApiData(data,cfg)
}
