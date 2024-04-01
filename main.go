package main

type command struct {
	name        string
	description string
	handler     func() error
}

func main() {
	startRepl()
}
