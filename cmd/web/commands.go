package main

import (
	"fmt"
	"os"

	"github.com/frankie-mur/pokedexcli/internal/models"
)

// Hold the next and previous urls for pokedex api requests
type Config struct {
	next  *string
	prev  *string
	cache *models.Cache
}

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
	//Check cache
	//val, exists := c.cache.Get(*c.next)
	//if exists {

	//	}
	data, err := models.GetTop20(c.next)
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
	data, err := models.GetTop20(c.prev)
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
