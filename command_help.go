package main

import (
	"fmt"
	"internal/pokecache"

	"github.com/nibi/pokedexcli/internal/pokedex"
)

func commandHelp(cfg *config, mem *pokecache.Cache, _ *pokedex.Pokedex, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n") // 2 \n are used - one in the func itself and 1 has been added
	commandMap := getCommands()
	for _, command := range commandMap {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}
