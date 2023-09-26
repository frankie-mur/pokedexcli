package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/frankie-mur/pokedexcli/internal/models"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display 20 locations of pokemon locations",
			callback:    commandMap,
		},
	}
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		msg := fmt.Sprintf("%s : %s", cmd.name, cmd.description)
		fmt.Println(msg)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandMap() error {
	data, err := models.GetTop20()
	if err != nil {
		fmt.Printf("An error occurred, %s\n", err.Error())
	}

	//Print top 20 names
	for _, res := range data.Results {
		fmt.Println(res.Name)
	}

	return nil
}
