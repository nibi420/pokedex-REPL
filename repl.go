package main

import (
	"internal/pokecache"
	"strings"

	"github.com/nibi/pokedexcli/internal/pokedex"
)

func cleanInput(text string) []string {

	text = strings.ToLower(text) // lowercase everything
	// text = strings.TrimSpace(text) // remove leading and trailing whitespace - might not be needed

	words := strings.Fields(text)

	return words

}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, mem *pokecache.Cache, pkdx *pokedex.Pokedex, arg string) error
}

type config struct {
	nextUrl string
	prevUrl string
}

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type specificLocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
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
			description: "Displays names of 20 locations areas in the Pokemon world. Each subsequent call to map displays the next 20 locations, and so on",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays names of 20 locations areas in the Pokemon world. Each subsequent call to map displays the previous 20 locations, and so on",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "See All Pokemon Located in an Area | needs an argument-> Area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Used to capture Pokemon | needs an argument-> Pokemon Name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Check the details of a pokemon within your Pokedex | needs an argument-> Pokemon Name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints all the pokemons captured by you",
			callback:    commandPokedex,
		},
	}

}
