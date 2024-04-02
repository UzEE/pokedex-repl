package main

type config struct {
	Next     *string
	Previous *string
}

type command struct {
	name        string
	description string
	handler     func(c *config) error
}
