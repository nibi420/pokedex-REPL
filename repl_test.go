package main

import (
	"fmt"
	"internal/pokecache"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{ // empty string input
			input:    "",
			expected: []string{},
		},
		{ // only whitespace input
			input:    "   ",
			expected: []string{},
		},
		{ // more whitespaces
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{ //uppercase lowercase test
			input:    "pIkAchu charizard blastOISE",
			expected: []string{"pikachu", "charizard", "blastoise"},
		},
		{ // whitespace, uppercase, lowercase test
			input:    "      zoobat  CHARmander       raichu   LUGIA",
			expected: []string{"zoobat", "charmander", "raichu", "lugia"},
		},
	}

	for _, c := range cases {

		actual := cleanInput(c.input)
		// lets check the length against the actual slice
		if len(actual) != len(c.expected) {
			t.Errorf("Actual: %v, Expected: %v", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Actual Word: %v, Expected Word: %v", word, expectedWord)
			}

			fmt.Printf("Passed %d of %d: Actual Word: %v, Expected Word: %v\n", i, len(actual), word, expectedWord)
		}
	}

}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area/",
			val: []byte(`{"count":1089,"next":"https://pokeapi.co/api/v2/location-area/?offset=20&limit=20","previous":null,"results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"},{"name":"eterna-city-area","url":"https://pokeapi.co/api/v2/location-area/2/"},{"name":"pastoria-city-area","url":"https://pokeapi.co/api/v2/location-area/3/"},{"name":"sunyshore-city-area","url":"https://pokeapi.co/api/v2/location-area/4/"},{"name":"sinnoh-pokemon-league-area","url":"https://pokeapi.co/api/v2/location-area/5/"},{"name":"oreburgh-mine-1f","url":"https://pokeapi.co/api/v2/location-area/6/"},{"name":"oreburgh-mine-b1f","url":"https://pokeapi.co/api/v2/location-area/7/"},{"name":"valley-windworks-area","url":"https://pokeapi.co/api/v2/location-area/8/"},{"name":"eterna-forest-area","url":"https://pokeapi.co/api/v2/location-area/9/"},{"name":"fuego-ironworks-area","url":"https://pokeapi.co/api/v2/location-area/10/"},{"name":"mt-coronet-1f-route-207","url":"https://pokeapi.co/api/v2/location-area/11/"},{"name":"mt-coronet-2f","url":"https://pokeapi.co/api/v2/location-area/12/"},{"name":"mt-coronet-3f","url":"https://pokeapi.co/api/v2/location-area/13/"},{"name":"mt-coronet-exterior-snowfall","url":"https://pokeapi.co/api/v2/location-area/14/"},{"name":"mt-coronet-exterior-blizzard","url":"https://pokeapi.co/api/v2/location-area/15/"},{"name":"mt-coronet-4f","url":"https://pokeapi.co/api/v2/location-area/16/"},{"name":"mt-coronet-4f-small-room","url":"https://pokeapi.co/api/v2/location-area/17/"},{"name":"mt-coronet-5f","url":"https://pokeapi.co/api/v2/location-area/18/"},{"name":"mt-coronet-6f","url":"https://pokeapi.co/api/v2/location-area/19/"},{"name":"mt-coronet-1f-from-exterior","url":"https://pokeapi.co/api/v2/location-area/20/"}]}`),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/?limit=20&offset=20",
			val: []byte(`{"count":1089,"next":"https://pokeapi.co/api/v2/location-area/?offset=40&limit=20","previous":"https://pokeapi.co/api/v2/location-area/?offset=0&limit=20","results":[{"name":"mt-coronet-1f-route-216","url":"https://pokeapi.co/api/v2/location-area/21/"},{"name":"mt-coronet-1f-route-211","url":"https://pokeapi.co/api/v2/location-area/22/"},{"name":"mt-coronet-b1f","url":"https://pokeapi.co/api/v2/location-area/23/"},{"name":"great-marsh-area-1","url":"https://pokeapi.co/api/v2/location-area/24/"},{"name":"great-marsh-area-2","url":"https://pokeapi.co/api/v2/location-area/25/"},{"name":"great-marsh-area-3","url":"https://pokeapi.co/api/v2/location-area/26/"},{"name":"great-marsh-area-4","url":"https://pokeapi.co/api/v2/location-area/27/"},{"name":"great-marsh-area-5","url":"https://pokeapi.co/api/v2/location-area/28/"},{"name":"great-marsh-area-6","url":"https://pokeapi.co/api/v2/location-area/29/"},{"name":"solaceon-ruins-2f","url":"https://pokeapi.co/api/v2/location-area/30/"},{"name":"solaceon-ruins-1f","url":"https://pokeapi.co/api/v2/location-area/31/"},{"name":"solaceon-ruins-b1f-a","url":"https://pokeapi.co/api/v2/location-area/32/"},{"name":"solaceon-ruins-b1f-b","url":"https://pokeapi.co/api/v2/location-area/33/"},{"name":"solaceon-ruins-b1f-c","url":"https://pokeapi.co/api/v2/location-area/34/"},{"name":"solaceon-ruins-b2f-a","url":"https://pokeapi.co/api/v2/location-area/35/"},{"name":"solaceon-ruins-b2f-b","url":"https://pokeapi.co/api/v2/location-area/36/"},{"name":"solaceon-ruins-b2f-c","url":"https://pokeapi.co/api/v2/location-area/37/"},{"name":"solaceon-ruins-b3f-a","url":"https://pokeapi.co/api/v2/location-area/38/"},{"name":"solaceon-ruins-b3f-b","url":"https://pokeapi.co/api/v2/location-area/39/"},{"name":"solaceon-ruins-b3f-c","url":"https://pokeapi.co/api/v2/location-area/40/"}]}`),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/limit=20&offset=80",
			val: []byte(`{"count":1089,"next":"https://pokeapi.co/api/v2/location-area/?offset=100&limit=20","previous":"https://pokeapi.co/api/v2/location-area/?offset=60&limit=20","results":[{"name":"iron-island-1f","url":"https://pokeapi.co/api/v2/location-area/120/"},{"name":"iron-island-b1f-left","url":"https://pokeapi.co/api/v2/location-area/121/"},{"name":"iron-island-b1f-right","url":"https://pokeapi.co/api/v2/location-area/122/"},{"name":"iron-island-b2f-right","url":"https://pokeapi.co/api/v2/location-area/123/"},{"name":"iron-island-b2f-left","url":"https://pokeapi.co/api/v2/location-area/124/"},{"name":"iron-island-b3f","url":"https://pokeapi.co/api/v2/location-area/125/"},{"name":"old-chateau-entrance","url":"https://pokeapi.co/api/v2/location-area/126/"},{"name":"old-chateau-dining-room","url":"https://pokeapi.co/api/v2/location-area/127/"},{"name":"old-chateau-2f-private-room","url":"https://pokeapi.co/api/v2/location-area/128/"},{"name":"old-chateau-2f","url":"https://pokeapi.co/api/v2/location-area/129/"},{"name":"old-chateau-2f-leftmost-room","url":"https://pokeapi.co/api/v2/location-area/130/"},{"name":"old-chateau-2f-left-room","url":"https://pokeapi.co/api/v2/location-area/131/"},{"name":"old-chateau-2f-middle-room","url":"https://pokeapi.co/api/v2/location-area/132/"},{"name":"old-chateau-2f-right-room","url":"https://pokeapi.co/api/v2/location-area/133/"},{"name":"old-chateau-2f-rightmost-room","url":"https://pokeapi.co/api/v2/location-area/134/"},{"name":"lake-verity-before-galactic-intervention","url":"https://pokeapi.co/api/v2/location-area/135/"},{"name":"lake-verity-after-galactic-intervention","url":"https://pokeapi.co/api/v2/location-area/136/"},{"name":"lake-valor-area","url":"https://pokeapi.co/api/v2/location-area/137/"},{"name":"lake-acuity-area","url":"https://pokeapi.co/api/v2/location-area/138/"},{"name":"valor-lakefront-area","url":"https://pokeapi.co/api/v2/location-area/139/"}]}`),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return

			}
		})

	}
}

func TestReadLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
