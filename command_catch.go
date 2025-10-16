package main

import (
	"math/rand"
	"fmt"
)

func commandCatch(cfg *config, pokemonName string) error {
	pokemon, err := cfg.pokeapiClient.FetchPokemon(pokemonName)
	if err != nil {
		return err
	}

	// chance of catching the pokemon by rolling a die
	// number of sides on the die scale with base exp
	dieSize := pokemon.BaseExperience / 10 + 5 // add 5 to avoid zero sides
	throw := rand.Intn(dieSize)
	target := dieSize - 10
	success := throw >= target	

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if !success {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
