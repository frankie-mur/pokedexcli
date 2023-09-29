package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/frankie-mur/pokedexcli/internal/models"
)

// Hold the next and previous urls for pokedex api requests
type Config struct {
	next    *string
	prev    *string
	cache   *models.Cache
	pokedex map[string]*models.Pokemon
}

var url = "https://pokeapi.co/api/v2/location"

func main() {
	//Initialize our scanner
	s := bufio.NewScanner(os.Stdin)

	prompt := "pokedex > "
	commands := GetCommands()

	config := &Config{
		next:    &url,
		prev:    nil,
		cache:   models.NewCache(10 * time.Second),
		pokedex: map[string]*models.Pokemon{},
	}

	for {
		fmt.Print(prompt)
		//Start our scanner
		s.Scan()
		//Get the text from scanner
		text := s.Text()
		input := cleanInput(text)

		command := input[0]
		var value string
		if len(input) > 1 {
			value = input[1]
		}

		//If the text is empty
		if len(command) == 0 {
			continue
		}
		//Check if command exists, if so execute the callback
		c, ok := commands[command]
		if ok {
			err := c.callback(config, value)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			//Else the command is not found
			fmt.Println("Command not found")
			continue
		}

	}
}
