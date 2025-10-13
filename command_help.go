package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for comm, desc := range supportedCommands() {
		fmt.Println(comm, desc.description)
	}
	fmt.Println()
	return nil
}
