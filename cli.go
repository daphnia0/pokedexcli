package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func printUnknown(text string) {
	fmt.Println(text, ": command not found")
	fmt.Println("")
}

func handleInvalidCmd(text string) {
	defer printUnknown(text)
}

func handleCmd(text string) {
	handleInvalidCmd(text)
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
func displayHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("map: goes to the next map")
	fmt.Println("mapb: goes back to the prev map")
	fmt.Println("clear: Clears the command table")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("")
	return nil
}

func clearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
