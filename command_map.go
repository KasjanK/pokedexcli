package main

import (
	"fmt"

	"github.com/KasjanK/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	locationsList, err := pokeapi.FetchLocationsList(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsList.Next
	cfg.prevLocationURL = locationsList.Previous

	for _, location := range locationsList.Results {
		fmt.Println(location.Name)
	}
	return nil
}
func commandMapb(cfg *config) error {
	if cfg.prevLocationURL == nil {

		fmt.Println("you're on the first page")
		return nil
	}
	locationsList, err := pokeapi.FetchLocationsList(cfg.prevLocationURL)
	if err != nil {
		return err 
	}

	cfg.nextLocationURL = locationsList.Next
	cfg.prevLocationURL = locationsList.Previous

	for _, location := range locationsList.Results {
		fmt.Println(location.Name)
	}
	return nil
}

