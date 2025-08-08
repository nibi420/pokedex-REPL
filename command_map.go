package main

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"io"
	"net/http"

	"github.com/nibi/pokedexcli/internal/pokedex"
)

func commandMap(cfg *config, mem *pokecache.Cache, _ *pokedex.Pokedex, _ string) error {
	client := &http.Client{} // make a client
	var data locationArea    // define a variable to get the response in a structured way
	url := "https://pokeapi.co/api/v2/location-area/"

	if cfg.nextUrl != "" {
		url = cfg.nextUrl
	}

	// here we would see if we already have the data in cache
	// if we have it:
	// we unmarshal the byte data into a Go Struct
	// in this case the data struct
	// print from there
	if value, ok := mem.Get(url); ok {

		if err := json.Unmarshal(value, &data); err != nil {
			return err
		}

		cfg.nextUrl = data.Next
		cfg.prevUrl = data.Previous

		for _, location := range data.Results {
			fmt.Println(location.Name)
		}
		return nil

	}

	// create a get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to Send GET Request")
		return err
	}
	// send a get request
	resp, err := client.Do(req)
	if err != nil {
		return err

	}

	// Let's add defer resp.Body.Close()
	defer resp.Body.Close()
	// decode the response Body
	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// since we have the data in bytes, lets add it to our cache
	mem.Add(url, byteData)

	if err := json.Unmarshal(byteData, &data); err != nil {
		return err
	}

	//set the next and previous url
	cfg.nextUrl = data.Next
	cfg.prevUrl = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil

}

func commandMapB(cfg *config, mem *pokecache.Cache, _ *pokedex.Pokedex, _ string) error {

	client := &http.Client{} // make a client
	var data locationArea    // define a variable to get the response in a structured way
	url := "https://pokeapi.co/api/v2/location-area/"

	// if we are *not* calling commandMapB the first time
	// then it means its a subsequent call
	// so we need to set our url to the prev url
	if cfg.prevUrl != "" {
		url = cfg.prevUrl
	}

	// here we would see if we already have the data in cache
	// if we have it:
	// we unmarshal the byte data into a Go Struct
	// in this case the data struct
	// print from there
	if value, ok := mem.Get(url); ok {

		if err := json.Unmarshal(value, &data); err != nil {
			return err
		}
		cfg.nextUrl = data.Next
		cfg.prevUrl = data.Previous

		for _, location := range data.Results {
			fmt.Println(location.Name)
		}
		return nil

	}

	// if we are here that means it was a cache miss
	// so now we need to make the http request
	// and also we need to keep that data in the cache

	// create a get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to Send GET Request")
		return err
	}
	// send a get request
	resp, err := client.Do(req)
	if err != nil {
		return err

	}
	// Let's add defer resp.Body.Close()
	defer resp.Body.Close()
	// convert the recieved the date to byte
	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// since we have the data in bytes, lets add it to our cache
	mem.Add(url, byteData)

	if err := json.Unmarshal(byteData, &data); err != nil {
		return err
	}

	//set the next and previous url
	cfg.nextUrl = data.Next
	cfg.prevUrl = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil

}
