// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliName string = "pokemoncli"

func main() {
	pokeMap := pokemonMap{Next: "https://pokeapi.co/api/v2/location-area/"}
	commands := map[string]interface{}{
		"help":  displayHelp,
		"map":   pokeMap.nextPokemonMap,
		"mapb":  pokeMap.prevPokemonMap,
		"clear": clearScreen,
	}
	reader := bufio.NewScanner(os.Stdin)

	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if command, exists := commands[text]; exists {
			command.(func() error)()
		} else if strings.EqualFold("exit", text) {
			return
		} else {
			handleCmd(text)
		}
		printPrompt()
	}
	fmt.Println()
}
