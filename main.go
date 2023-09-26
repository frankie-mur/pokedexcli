package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	prompt := "pokedex >"
	commands := GetCommands()

	for {
		fmt.Print(prompt)
		s.Scan()
		text := (s.Text())
		command := cleanInput(text)[0]

		commands[command].callback()

	}
}
