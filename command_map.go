package main

import (
	"errors"
	"fmt"

	"github.com/UzEE/pokedexcli/internal/api"
)

func mapCommand(c *config) error {
	locations, err := makeRequest(c.Next, *c.client)
	if err != nil {
		return err
	}

	handleResponse(locations)
	updatePages(c, locations)
	return nil
}

func mapBCommand(c *config) error {
	if c.Previous == nil {
		return errors.New("no previous locations. Use \"map\" command to see the next locations")
	}

	locations, err := makeRequest(c.Previous, *c.client)
	if err != nil {
		return err
	}

	handleResponse(locations)
	updatePages(c, locations)
	return nil
}

func makeRequest(url *string, client api.Client) (api.PagedResourceList, error) {
	fmt.Printf("Fetching...")
	defer fmt.Printf("\033[2K\r")

	return client.ListLocationArea(url)
}

func handleResponse(list api.PagedResourceList) {
	for _, loc := range list.Results {
		fmt.Println(loc.Name)
	}
}

func updatePages(c *config, list api.PagedResourceList) {
	if list.Next != nil {
		next := *list.Next
		c.Next = &next
	} else {
		c.Next = nil
	}

	if list.Previous != nil {
		previous := *list.Previous
		c.Previous = &previous
	} else {
		c.Previous = nil
	}
}
