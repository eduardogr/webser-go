package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	server, err := InitializeServer()

	if err != nil {
		return err
	}

	err = server.HandleRequests()

	if err != nil {
		return err
	}

	return nil
}
