package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/frankie-mur/pokedexcli/internal/models"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
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
		"catch": {
			name:        "catch",
			description: "Used to catch a specific pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Used to inspect a specific pokemon (if in pokedex)",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all of the pokemon in pokedex",
			callback:    commandPokedex,
		},
	}
}

func commandHelp(c *Config, val string) error {
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

func commandExit(c *Config, val string) error {
	os.Exit(0)
	return nil
}

func commandMap(c *Config, val string) error {
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

func commandMapb(c *Config, val string) error {
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

func commandExplore(c *Config, val string) error {
	fmt.Println("Exploring city....")
	data, err := models.GetNamesFromArea(c.cache, val)
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

func commandCatch(c *Config, val string) error {
	//Check pokedex for pokemon
	_, ok := c.pokedex[val]
	if ok {
		return errors.New(val + " is already caught!")
	}
	data, err := models.CatchPokemon(val)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokemon at %s...\n", val)
	didCatch := tryCatch(data.BaseExperience)
	if !didCatch {
		return errors.New("missed the Pokemon! maybe try again")
	}

	//Store the pokemon in users pokedex
	c.pokedex[val] = data
	fmt.Printf("Caught %s and stored in pokedex!\n", val)

	return nil
}

func commandInspect(c *Config, val string) error {
	if val == "" {
		return errors.New("enter a valid pokemon")
	}

	//Check the pokedex
	pokemon, ok := c.pokedex[val]
	if !ok {
		return errors.New("pokemon is not in pokedex")
	}

	//Print the pokemon's stats
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, c := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", c.Stat.Name, c.BaseStat)
	}
	fmt.Println("Types:")
	for _, c := range pokemon.Types {
		fmt.Printf(" - %s\n", c.Type.Name)
	}

	return nil
}

func commandPokedex(c *Config, val string) error {
	fmt.Println("Your Pokedex:")

	for k := range c.pokedex {
		fmt.Printf(" - %s\n", k)
	}

	return nil
}
