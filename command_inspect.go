package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n",stat.Stat.Name, stat.BaseStat)

	}
	fmt.Println("Types:")	
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  -%s\n", typeInfo.Type.Name)
	}
	return nil
}
