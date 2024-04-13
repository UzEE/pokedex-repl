package main

import (
	"fmt"
	"math/rand"
)

func inspectCommand(c *config, args ...string) error {
	if len(args) == 0 && c.latestCaught != nil {
		args = []string{*c.latestCaught}
	}

	if len(args) != 1 {
		return fmt.Errorf("missing required argument: <pokemon name>")
	}

	name := args[0]

	entry, ok := c.pokedex[name]
	if !ok {
		return fmt.Errorf("you haven't caught %s", name)
	}

	feet, inches := decimetersToFeetAndInches(entry.Height)

	printLine("No. %d, %s", entry.SpeciesEntry.ID, getFormattedName(c, entry))

	for _, genus := range entry.SpeciesEntry.Genera {
		if genus.Language.Name == "en" {
			printLine("%s", genus.Genus)
			break
		}
	}

	flavors := make([]string, 0, len(entry.SpeciesEntry.FlavorTextEntries))

	for _, flavor := range entry.SpeciesEntry.FlavorTextEntries {
		if flavor.Language.Name == "en" {
			flavors = append(flavors, flavor.FlavorText)
		}
	}

	printLine("%s", stripBreaks(flavors[rand.Int()%len(flavors)]))

	printLine("Height: %d' %d\"", feet, inches)

	if entry.Weight >= 100 {
		printLine("Weight: %.2f kg", float64(entry.Weight)/100)
	} else {
		printLine("Weight: %d g", entry.Weight*10)
	}

	printLine("Abilities:")
	for _, ability := range entry.Abilities {
		printLine(" - %s", ability.Ability.Name)
	}

	printLine("Types:")
	for _, t := range entry.Types {
		printLine(" - %s", t.Type.Name)
	}

	printLine("Stats:")
	for _, stat := range entry.Stats {
		printLine(" - %s: %d", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
