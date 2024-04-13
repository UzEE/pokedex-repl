package main

import "fmt"

func exploreCommand(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected location name in 1 argument, got %d", len(args))
	}

	name := args[0]

	printLine("Exploring %s...", name)

	loc, err := c.client.GetLocationDetails(name)
	if err != nil {
		return err
	}

	if len(loc.Areas) == 0 {
		printLine("No areas found.")
		return nil
	} else {
		printLine("Areas:")
	}

	for _, area := range loc.Areas {
		printLine(" - %s", area.Name)
	}

	c.currentLocation = &loc

	if c.currentRegion == nil {
		region, err := c.client.GetRegionDetails(loc.Region.Name)
		if err != nil {
			return err
		}

		c.currentRegion = &region
	}

	c.prompt = fmt.Sprintf("Pokédex [in %s @ %s]> ", loc.Region.Name, loc.Name)

	printLine()
	printLine("Type:")
	printLine("  %sscan <area>%s to PokéScan the wild Pokémon in the area", c.colors.Blue, c.colors.Reset)
	printLine("  %sexplore <location>%s to explore a different location", c.colors.Blue, c.colors.Reset)
	printLine("  %svisit <region>%s to visit a different region", c.colors.Blue, c.colors.Reset)

	return nil
}
