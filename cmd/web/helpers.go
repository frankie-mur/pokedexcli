package main

import (
	"strings"

	"github.com/frankie-mur/pokedexcli/internal/models"
	"golang.org/x/exp/rand"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func setConfig(c *Config, d *models.PokeDexLocation) {
	c.next = d.Next
	c.prev = d.Previous
}

func tryCatch(exp int) bool {
	//Checks if a random number between 0 and exp is less than 50
	//May want to change this in future as not sure of exp values
	return rand.Intn(exp) < 50
}
