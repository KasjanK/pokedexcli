package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, parameter string) error {
	locationsList, err := cfg.pokeapiClient.FetchLocationsList(cfg.nextLocationURL)
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
func commandMapb(cfg *config, parameter string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationsList, err := cfg.pokeapiClient.FetchLocationsList(cfg.prevLocationURL)
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

