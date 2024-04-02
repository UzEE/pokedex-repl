package main

import "os"

func exitCommand(_ *config) error {
	os.Exit(0)
	return nil
}
