package main

func expCommand(c *config, _ ...string) error {
	area, err := c.client.GetLocationAreaDetails("valley-windworks-area")
	if err != nil {
		return err
	}

	printLine(area)

	return nil
}
