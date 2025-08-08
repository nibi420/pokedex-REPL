package main

import (
	"bufio"
	"fmt"
	"internal/pokecache"
	"os"
	"time"

	"github.com/nibi/pokedexcli/internal/pokedex"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var configuration config
	interval := 60 * time.Second
	cache := pokecache.NewCache(interval) // this is a pointer to the cache
	mainPokedex := pokedex.NewPokedex()
	for {

		fmt.Printf("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		if len(input) == 0 {
			continue
		}

		cleanedInput := cleanInput(input) // cleanedInput is an array
		commandMap := getCommands()

		for i, word := range cleanedInput {
			if command, ok := commandMap[word]; ok {
				arg := ""                    // the argument is empty
				if i+1 < len(cleanedInput) { // so the word after the command is the argument
					arg = cleanedInput[i+1]
				}
				command.callback(&configuration, cache, mainPokedex, arg)

			}
		}

	}

}
