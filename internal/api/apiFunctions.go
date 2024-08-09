package apiFunctions

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type PokemonMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetApiData(endpoint string) []byte {
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

func UnmarshalApi(body []byte, pokeMap *PokemonMap) {
	errUnmarshall := json.Unmarshal(body, &pokeMap)
	if errUnmarshall != nil {
		log.Fatal(errUnmarshall)
	}
}
func (pokeMap *PokemonMap) PrintMap() {
	for _, v := range pokeMap.Results {
		println(v.Name)
	}
}
func (pokeMap *PokemonMap) NextPokemonMap() error {
	if pokeMap.Next != "" {
		body := GetApiData(pokeMap.Next)
		UnmarshalApi(body, pokeMap)
		pokeMap.PrintMap()
		return nil
	} else {
		return errors.New("no preveous 20 map found")
	}
}

func (pokeMap *PokemonMap) PrevPokemonMap() error {
	if pokeMap.Previous != "" {
		body := GetApiData(pokeMap.Previous)
		UnmarshalApi(body, pokeMap)
		pokeMap.PrintMap()
		return nil
	} else {
		println("no preveous 20 map found")
		return errors.New("no preveous 20 map found")
	}
}
