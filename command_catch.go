package main

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"io"
	"math"
	"math/rand"
	"net/http"

	"github.com/nibi/pokedexcli/internal/pokedex"
)

func commandCatch(cfg *config, mem *pokecache.Cache, pkdx *pokedex.Pokedex, arg string) error {

	client := &http.Client{}
	var data pokedex.Pokemon
	caught := false

	url := "https://pokeapi.co/api/v2/pokemon/" + arg // this assumes arg is a name only

	if value, ok := mem.Get(url); ok { // this gives us the data if we already have the pokemon structure cached

		if err := json.Unmarshal(value, &data); err != nil {
			return err
		}

		// for now we are not checking if we already have the pokemon caught, we just catch it

		fmt.Printf("Throwing a Pokeball at %v...\n", data.Name)

		// now lets calculate the probability when we throw the pokeball

		caught = decreasingChance(data.BaseExperience)

		if caught {
			pkdx.AddToPokedex(data.Name, data)
			fmt.Printf("%v was caught\n", data.Name)
		} else {
			fmt.Printf("%v escaped \n", data.Name)
		}

		return nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to Send Get Request")
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to recieve response")
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

	// now we need to actually throw the pokeball

	fmt.Printf("Throwing a Pokeball at %v...\n", data.Name)

	// now lets calculate the probability when we throw the pokeball

	caught = decreasingChance(data.BaseExperience)

	if caught {
		pkdx.AddToPokedex(data.Name, data)
		fmt.Printf("%v was caught\n", data.Name)
	} else {
		fmt.Printf("%v escaped \n", data.Name)
	}

	return nil

}

func decreasingChance(x int) bool { // this actually decides if the pokemon shall be captured or not

	if x < 0 {
		x = 0
	}

	decay := 0.0045

	chance := math.Exp(-decay * float64(x))

	r := rand.Float64()

	return r < chance

}
