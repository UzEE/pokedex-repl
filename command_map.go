package main

import (
	"errors"
	"fmt"

	"github.com/UzEE/pokedexcli/internal/api"
	"github.com/UzEE/pokedexcli/internal/api/types"
)

func mapCommand(c *config, _ ...string) error {
	locations, err := makeRequest(c.NextArea, *c.client)
	if err != nil {
		return err
	}

	handleResponse(locations)
	updatePages(c, locations)
	return nil
}

func mapBCommand(c *config, _ ...string) error {
	if c.PreviousArea == nil {
		return errors.New("no previous locations. Use \"map\" command to see the next locations")
	}

	locations, err := makeRequest(c.PreviousArea, *c.client)
	if err != nil {
		return err
	}

	handleResponse(locations)
	updatePages(c, locations)
	return nil
}

func makeRequest(url *string, client api.Client) (types.PagedResourceList, error) {
	fmt.Printf("Fetching...")
	defer fmt.Printf("\033[2K\r")

	return client.ListLocationAreas(url)
}

func handleResponse(list types.PagedResourceList) {
	for _, loc := range list.Results {
		printLine(loc.Name)
	}
}

func updatePages(c *config, list types.PagedResourceList) {
	if list.Next != nil {
		next := *list.Next
		c.NextArea = &next
	} else {
		c.NextArea = nil
	}

	if list.Previous != nil {
		previous := *list.Previous
		c.PreviousArea = &previous
	} else {
		c.PreviousArea = nil
	}
}
