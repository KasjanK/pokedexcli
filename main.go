package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		fmt.Println("Your command was:", strings.ToLower(input[0]))

	}
}

func cleanInput(text string) []string {
	textSlice := strings.ToLower(text)
	words := strings.Fields(textSlice)
	return words
}
