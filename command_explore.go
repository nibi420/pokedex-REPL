package main

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"io"
	"net/http"

	"github.com/nibi/pokedexcli/internal/pokedex"
)

func commandExplore(_ *config, mem *pokecache.Cache, _ *pokedex.Pokedex, arg string) error {

	// Implementation
	/*

		take the argument
		see what location it is -> 	pull up pokemons accordingly
		[
			adding the argument at the end of the url
			(It can be name or id)
			and sending it over as an http request
			if the http request fails and the location area is not valid?
			nothing, because the API, takes care of that
			now in the results, we just make a for loop to print out the pokemons, as easy as that.

		]
	*/

	client := &http.Client{}
	var data specificLocationArea // define a variable to get the response in a structured way
	url := "https://pokeapi.co/api/v2/location-area/" + arg

	if value, ok := mem.Get(url); ok {

		if err := json.Unmarshal(value, &data); err != nil {
			return err
		}

		for _, pokemonDetails := range data.PokemonEncounters {
			fmt.Println(pokemonDetails.Pokemon.Name)
		}
		return nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to Send GET Request")
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to get Response")
		return err
	}
	defer resp.Body.Close()

	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// since we have the data in bytes, lets add it to our cache
	mem.Add(url, byteData)

	if err := json.Unmarshal(byteData, &data); err != nil {
		return err
	}

	for _, pokemonDetails := range data.PokemonEncounters {
		fmt.Println(pokemonDetails.Pokemon.Name)
	}
	return nil

}
