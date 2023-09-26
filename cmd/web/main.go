package main

import (
	"bufio"
	"fmt"
	"os"
)

var url = "https://pokeapi.co/api/v2/location"

func main() {
	//Initialize our scannr
	s := bufio.NewScanner(os.Stdin)

	prompt := "pokedex > "
	commands := GetCommands()

	config := &Config{
		next: &url,
		prev: nil,
	}

	for {
		fmt.Print(prompt)
		//Start our scanner
		s.Scan()
		//Get the text from scanner
		text := s.Text()
		command := cleanInput(text)[0]
		//If the text is empty
		if len(command) == 0 {
			continue
		}
		//Check if command exists, if so execute the callback
		c, ok := commands[command]
		if ok {
			err := c.callback(config)
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
