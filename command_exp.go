package main

import (
	"fmt"
)

func expCommand(c *config, _ ...string) error {
	area, err := c.client.GetLocationAreaDetails("valley-windworks-area")
	if err != nil {
		return err
	}

	fmt.Println(area)

	return nil
}
