# Pokédex CLI 🕹️

A simple interactive command-line Pokédex that lets you explore the Pokémon world via the [PokéAPI](https://pokeapi.co/)!  
Catch Pokémon, explore locations, and inspect your growing Pokédex — all from the terminal.

---

## ✨ Features

- 🔍 **Explore** different location areas to see what Pokémon appear there  
- 🗺️ **Map** navigation (next/previous) through paginated location areas  
- 🎯 **Catch** Pokémon by name, with success based on their base experience  
- 📖 **Inspect** caught Pokémon to view their stats, type, and more  
- 📘 **Pokedex** command to list all Pokémon you've successfully caught  
- ⚡ **Caching** system to avoid unnecessary API calls (with timed expiration)

---

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or later installed on your system

### Installation

1. Clone the repository:

```bash
git clone https://github.com/your-username/pokedexcli.git
cd pokedexcli
```
2. Run the Application 
```bash
go run .
```

### Available Commands 
help                          # List available commands
exit                          # Exit the Pokédex CLI
map                           # Show the next page of location areas
mapb                          # Show the previous page of location areas
explore <location_name>       # View Pokémon encounters in a location
catch <pokemon_name>          # Attempt to catch a Pokémon
inspect <pokemon_name>        # View stats for a caught Pokémon
pokedex                       # Show all caught Pokémon

### Catch Mechanics 
When you run:
```bash
catch pikachu
```
The catch success is determined randomly, but weighted by the Pokémon's base experience — the stronger the Pokémon, the harder it is to catch.

The catch formula looks like this:
```bash
catchChance := 0.9 - (baseExperience / 1000)
```
A random float [0.0 - 1.0] is rolled. If the roll is less than catchChance, the Pokémon is caught!