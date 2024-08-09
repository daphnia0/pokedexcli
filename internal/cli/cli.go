package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PrintPrompt(cliName string) {
	fmt.Print(cliName, " > ")
}

func PrintUnknown(text string) {
	fmt.Println(text, ": command not found")
	fmt.Println("")
}

func HandleInvalidCmd(text string) {
	defer PrintUnknown(text)
}

func HandleCmd(text string) {
	HandleInvalidCmd(text)
}

func CleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
func DisplayHelp() error {
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

func ClearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
