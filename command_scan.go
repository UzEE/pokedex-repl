package main

import "fmt"

func scanCommand(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected area name in 1 argument, got %d", len(args))
	}

	name := args[0]

	printLine("Running PokéScan...")

	area, err := c.client.GetLocationAreaDetails(name)
	if err != nil {
		return err
	}

	if len(area.PokemonEncounters) == 0 {
		printLine("No Pokémon encounters found.")
		return nil
	} else {
		printLine("Found Pokémon:")
	}

	count := len(area.PokemonEncounters)
	detail_chan := make(chan string, count)

	for _, encounter := range area.PokemonEncounters {
		maxChance := 0

		for _, version := range encounter.VersionDetails {
			if version.MaxChance > maxChance {
				maxChance = version.MaxChance
			}
		}

		go func() {
			mon, err := c.client.GetPokemonAndSpecies(encounter.Pokemon.Name)
			if err != nil {
				detail_chan <- fmt.Sprintf("Error fetching Pokémon details: %v", encounter.Pokemon.Name)
				return
			}

			formattedName := getFormattedName(c, mon)

			detail_chan <- fmt.Sprintf(" - %s (chance: %d)", formattedName, maxChance)
		}()
	}

	for i := 0; i < count; i++ {
		printLine(<-detail_chan)
	}

	close(detail_chan)

	c.currentArea = &area

	if c.currentLocation == nil {
		loc, err := c.client.GetLocationDetails(area.Location.Name)
		if err != nil {
			return err
		}

		c.currentLocation = &loc
	}

	if c.currentRegion == nil {
		region, err := c.client.GetRegionDetails(c.currentLocation.Region.Name)
		if err != nil {
			return err
		}

		c.currentRegion = &region
	}

	c.prompt = fmt.Sprintf("Pokédex [in %s @ %s | %s]> ", c.currentRegion.Name, area.Location.Name, area.Name)

	printLine()
	printLine("Type:")
	printLine("  %sencounter%s to try and encounter a random Pokémon", c.colors.Blue, c.colors.Reset)
	printLine("  %sexplore <location>%s to explore a different location", c.colors.Blue, c.colors.Reset)

	return nil
}
