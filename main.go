package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name 			string
	description 	string
	callback 		func(*config) error
}

type pokedexLocation struct {
	Next     string `json:"next"`
	Previous string	`json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type config struct {
	nextLocationURL string
	prevLocationURL string
}

func fetchLocationsList(url string) (pokedexLocation, error) {
	res, err := http.Get(url)
	if err != nil {
		return pokedexLocation{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return pokedexLocation{}, fmt.Errorf("Error reading response body: %w", err)
	}
	locationsList := pokedexLocation{}
	err = json.Unmarshal(body, &locationsList)
	if err != nil {
		return pokedexLocation{}, err
	}
	return locationsList, nil
}

func commandMap(cfg *config) error {
	url := cfg.nextLocationURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	locationsList, err :=  fetchLocationsList(url)
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
	if cfg.prevLocationURL == "" {
		
		fmt.Println("you're on the first page")
		return nil
	}
	locationsList, err :=  fetchLocationsList(cfg.prevLocationURL)
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{ nextLocationURL: "https://pokeapi.co/api/v2/location-area/" }

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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

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

func cleanInput(text string) []string {
	textSlice := strings.ToLower(text)
	words := strings.Fields(textSlice)
	return words
}
