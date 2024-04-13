package main

func exitCommand(_ *config, _ ...string) error {
	return ErrExitSafe
}
