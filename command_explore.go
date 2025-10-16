package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, locationArea string) error {
	if locationArea == "" {
		return errors.New("no area found, try again")
	}

	location, err := cfg.pokeapiClient.FetchLocation(locationArea)
	if err != nil {
		return nil
	}

	fmt.Printf("Exploring %s...\n", locationArea)
	fmt.Println("Found Pokemon:")
	for _, name := range location.PokemonEncounters {
		fmt.Println(" - " + name.Pokemon.Name)
	}
	return nil
}
