package main

import (
	"strings"

	"github.com/frankie-mur/pokedexcli/internal/models"
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
