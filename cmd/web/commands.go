package main

import (
	"fmt"
	"os"

	"github.com/frankie-mur/pokedexcli/internal/models"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 pokemon locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display the name of pokemon in a specific location",
			callback:    commandExplore,
		},
	}
}

func commandHelp(c *Config) error {
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

func commandExit(c *Config) error {
	os.Exit(0)
	return nil
}

func commandMap(c *Config) error {
	data, err := models.GetTop20(c.cache, c.next)
	if err != nil {
		return err
	}

	//Print top 20 names
	for _, res := range data.Results {
		fmt.Println(res.Name)
	}

	setConfig(c, data)

	return nil
}

func commandMapb(c *Config) error {
	data, err := models.GetTop20(c.cache, c.prev)
	if err != nil {
		return err
	}

	//Print top 20 names
	for _, res := range data.Results {
		fmt.Println(res.Name)
	}

	setConfig(c, data)

	return nil
}

// TODO: Add name as a parameter and search by name instead of hard code
func commandExplore(c *Config) error {
	fmt.Println("Exploring city....")
	data, err := models.GetNamesFromArea(c.cache, "pastoria-city-area")
	if err != nil {
		return err
	}
	fmt.Printf("Found Pokemon:\n")
	for _, res := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", res.Pokemon.Name)
	}

	//fmt.Println(data)
	return nil
}
