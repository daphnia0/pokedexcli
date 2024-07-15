package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type pokemonMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getApiData(endpoint string) []byte {
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return body
}

func unmarshalApi(body []byte, pokeMap *pokemonMap) {
	errUnmarshall := json.Unmarshal(body, &pokeMap)
	if errUnmarshall != nil {
		log.Fatal(errUnmarshall)
	}
}
func (pokeMap *pokemonMap) printMap() {
	for _, v := range pokeMap.Results {
		println(v.Name)
	}
}
func (pokeMap *pokemonMap) nextPokemonMap() error {
	if pokeMap.Next != "" {
		body := getApiData(pokeMap.Next)
		unmarshalApi(body, pokeMap)
		pokeMap.printMap()
		return nil
	} else {
		return errors.New("no preveous 20 map found")
	}
}

func (pokeMap *pokemonMap) prevPokemonMap() error {
	if pokeMap.Previous != "" {
		body := getApiData(pokeMap.Previous)
		unmarshalApi(body, pokeMap)
		pokeMap.printMap()
		return nil
	} else {
		println("no preveous 20 map found")
		return errors.New("no preveous 20 map found")
	}
}
