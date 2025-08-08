package main

import (
	"fmt"
	"internal/pokecache"

	"github.com/nibi/pokedexcli/internal/pokedex"
)

func commandInspect(cfg *config, mem *pokecache.Cache, pkdx *pokedex.Pokedex, arg string) error {

	val, ok := pkdx.GetPokemon(arg)
	if !ok {
		fmt.Println("??? — You haven't caught this Pokémon yet. Its details are shrouded in mystery.")
	} else {
		printPokemonDetails(val)
	}
	return nil

}

func printPokemonDetails(p pokedex.Pokemon) {
	fmt.Println("📘 Pokedex Entry")
	fmt.Println("---------------")
	fmt.Println("Name:", p.Name)
	fmt.Println("ID:", p.ID)
	fmt.Println("Species:", p.Species.Name)
	fmt.Println("Base Experience:", p.BaseExperience)
	fmt.Println("Height:", p.Height)
	fmt.Println("Weight:", p.Weight)

	fmt.Println("\nTypes:")
	for _, t := range p.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}

	fmt.Println("\nAbilities:")
	for _, a := range p.Abilities {
		hidden := ""
		if a.IsHidden {
			hidden = " (Hidden)"
		}
		fmt.Printf("- %s%s\n", a.Ability.Name, hidden)
	}

	fmt.Println("\nSprite:")
	fmt.Println(p.Sprites.FrontDefault)
}
