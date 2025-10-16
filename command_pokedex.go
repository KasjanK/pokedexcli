package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, parameter string) error {
	pokedex := cfg.caughtPokemon
	if len(pokedex) == 0 {
		return errors.New("Your Pokedex in empty!")
	}
	fmt.Println("Your pokedex:")
	for _, pokemon := range pokedex {
		fmt.Printf("  -%s\n", pokemon.Name)
	}
	return nil
}
