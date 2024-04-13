package main

import (
	"fmt"

	"github.com/UzEE/pokedexcli/internal/api/types/location"
)

func visitCommand(c *config, args ...string) error {
	if len(args) == 0 {
		if c.currentRegion == nil {
			return fmt.Errorf("no region name in argument, and no current region set")
		}
	} else if len(args) != 1 {
		return fmt.Errorf("expected region name in 1 argument, got %d", len(args))
	} else if len(args) == 1 {

		if c.currentRegion != nil {
			if args[0] == "prev" {

				if c.LocationPage == 1 {
					return fmt.Errorf("already at the first page")
				}

				c.LocationPage -= 2
			} else if args[0] != c.currentRegion.Name {
				c.currentRegion = nil
				c.LocationPage = 0
			} else {
				c.LocationPage = 0
			}
		}
	}

	var region location.Region

	if c.currentRegion != nil {
		region = *c.currentRegion
	} else {
		name := args[0]
		var err error

		region, err = c.client.GetRegionDetails(name)
		if err != nil {
			return err
		}

		c.currentRegion = &region
		printLine("Visiting %s", region.Name)
	}

	indexLimit := min((c.LocationPage+1)*c.LocationPageSize, len(region.Locations))

	printLine("Locations (page %d of %d): ", c.LocationPage+1, len(region.Locations)/c.LocationPageSize+1)

	for i := (c.LocationPage * c.LocationPageSize); i < indexLimit; i++ {
		location := region.Locations[i]
		printLine(" - %s", location.Name)
	}

	c.LocationPage += 1
	c.prompt = fmt.Sprintf("Pokédex [in %s]> ", region.Name)

	printLine()
	printLine("Type:")
	printLine("  %svisit%s to view next page", c.colors.Blue, c.colors.Reset)
	printLine("  %svisit prev%s to go back to the previous page", c.colors.Blue, c.colors.Reset)
	printLine("  %svisit <region>%s to visit a different region", c.colors.Blue, c.colors.Reset)
	printLine("  %sexplore <location>%s to explore a location", c.colors.Blue, c.colors.Reset)

	return nil

}
