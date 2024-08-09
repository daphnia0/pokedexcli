// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	api "github.com/daphnia0/pokedexcli/internal/api"
	cli "github.com/daphnia0/pokedexcli/internal/cli"
)

var cliName string = "pokemoncli"

func main() {
	pokeMap := api.PokemonMap{Next: "https://pokeapi.co/api/v2/location-area/"}
	commands := map[string]interface{}{
		"help":  cli.DisplayHelp,
		"map":   pokeMap.NextPokemonMap,
		"mapb":  pokeMap.PrevPokemonMap,
		"clear": cli.ClearScreen,
	}
	reader := bufio.NewScanner(os.Stdin)

	cli.PrintPrompt(cliName)
	for reader.Scan() {
		text := cli.CleanInput(reader.Text())
		if command, exists := commands[text]; exists {
			command.(func() error)()
		} else if strings.EqualFold("exit", text) {
			return
		} else {
			cli.HandleCmd(text)
		}
		cli.PrintPrompt(cliName)
	}
	fmt.Println()
}
