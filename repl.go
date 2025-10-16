package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KasjanK/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name 			string
	description 	string
	callback 		func(*config, string) error
}

type config struct {
	pokeapiClient 	pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon	map[string]pokeapi.Pokemon
}

func supportedCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Shows 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Shows previous 20 location areas",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Shows all pokemon encounters in a specific area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Make an attempt to catch a pokemon",
			callback: commandCatch,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		var parameter string
		if len(input) > 1 {
			parameter = input[1]
		}

		commandName := input[0]
		command, ok := supportedCommands()[commandName]
		if ok { 
			err := command.callback(cfg, parameter)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	textSlice := strings.ToLower(text)
	words := strings.Fields(textSlice)
	return words
}
