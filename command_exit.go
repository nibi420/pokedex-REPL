package main

import (
	"fmt"
	"internal/pokecache"
	"os"

	"github.com/nibi/pokedexcli/internal/pokedex"
)

func commandExit(cfg *config, mem *pokecache.Cache, _ *pokedex.Pokedex, _ string) error {

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}
