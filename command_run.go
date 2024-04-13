package main

func runCommand(c *config, _ ...string) error {
	if c.currentEncounter == nil {
		printLine("You are not in any encounter.")
		return nil
	}

	pokemon, err := c.client.GetPokemonAndSpecies(c.currentEncounter.Name)
	if err != nil {
		return err
	}

	printLine("You ran away from the wild %s!", getFormattedName(c, pokemon))

	c.currentEncounter = nil
	return nil
}
