package main

import (
	"time"
	"strings"
)

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
func createPokemon(p pokemonData) Pokemon {
	types := []string{}
	for _, ty := range p.Types {
		types = append(types, ty.Type.Name)
	}

	var statistics Stats

	for _, stat := range p.Stats {
		switch stat.Stat.Name {
		case "hp":
			statistics.hp = stat.BaseStat
		case "attack":
			statistics.attack = stat.BaseStat
		case "defense":
			statistics.defense = stat.BaseStat
		case "special-attack":
			statistics.specialAttack = stat.BaseStat
		case "special-defense":
			statistics.specialDefense = stat.BaseStat
		case "speed":
			statistics.speed = stat.BaseStat
		}
	}

	return Pokemon{
		Name:           p.Name,
		BaseExperience: p.BaseExperience,
		Height:         p.Height,
		Weight:         p.Weight,
		Types:          types,
		Statistics:     statistics,
		CaughtAt:       time.Now(),
	}

}
