package main

import (
	"github.com/UzEE/pokedexcli/internal/api"
)

type config struct {
	client *api.Client

	Next     *string
	Previous *string
}

type command struct {
	name        string
	description string
	handler     func(c *config) error
}
