package main

import "os"

func exitCommand(_ *config, _ ...string) error {
	os.Exit(0)
	return nil
}
