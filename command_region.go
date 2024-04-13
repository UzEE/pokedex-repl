package main

import "fmt"

func regionCommand(c *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("expected 0 arguments, got %d", len(args))
	}

	regions, err := c.client.ListRegions()
	if err != nil {
		return err
	}

	printLine("Showing %d of %d regions:", len(regions.Results), regions.Count)

	for _, region := range regions.Results {
		printLine(region.Name)
	}

	if regions.Next != nil {
		next := *regions.Next
		c.NextRegion = &next
	} else {
		c.NextRegion = nil
	}

	if regions.Previous != nil {
		previous := *regions.Previous
		c.PreviousRegion = &previous
	} else {
		c.PreviousRegion = nil
	}

	printLine()
	printLine("Type:")
	printLine("  %svisit <region>%s to visit a specific region", c.colors.Blue, c.colors.Reset)

	return nil
}
