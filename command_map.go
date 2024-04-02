package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type NamedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func mapCommand(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if c.Next != nil {
		url = *c.Next
	} else {
		fmt.Println("No next locations. Use \"mapb\" command to see the previous locations.")
		return nil
	}

	return handleMapCommand(url, c)
}

func mapBCommand(c *config) error {
	var url string

	if c.Previous != nil {
		url = *c.Previous
	} else {
		fmt.Println("No previous locations. Use \"map\" command to see the next locations.")
		return nil
	}

	return handleMapCommand(url, c)
}

func handleMapCommand(url string, c *config) error {
	fmt.Printf("Loading...")
	res, err := http.Get(url)
	fmt.Printf("\033[2K\r")

	if err != nil {
		log.Printf("Failed to fetch map data: %v", err)
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode >= 400 {
		log.Printf("Failed to fetch map data: %s", body)
		return fmt.Errorf("failed to fetch map data. Status Code: %d", res.StatusCode)
	}

	if err != nil {
		log.Printf("Failed to read map data: %v", err)
		return err
	}

	list := PagedResourceList{}
	err = json.Unmarshal(body, &list)

	if err != nil {
		log.Printf("Failed to unmarshal map data: %v", err)
		return err
	}

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

	for _, loc := range list.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
