package main

import (
	"fmt"
	"internal/pokecache"

	"github.com/nibi/pokedexcli/internal/pokedex"
)

func commandPokedex(cfg *config, mem *pokecache.Cache, pkdx *pokedex.Pokedex, _ string) error {

	if len(pkdx.Directory) == 0 {
		fmt.Println("Professor Oak: It looks like you haven’t caught any Pokémon yet!")
		return nil
	}
	for pokemon := range pkdx.Directory {
		fmt.Println(pokemon)
	}
	return nil

}
