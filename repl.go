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
	callback 		func(*config) error
}

type config struct {
	pokeapiClient 	pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
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
		commandName := input[0]

		command, ok := supportedCommands()[commandName]
		if ok { 
			err := command.callback(cfg)
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
