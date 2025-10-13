package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLocationsList(pageURL *string) (PokedexLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return PokedexLocation{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokedexLocation{}, fmt.Errorf("Error reading response body: %w", err)
	}
	locationsList := PokedexLocation{}
	err = json.Unmarshal(body, &locationsList)
	if err != nil {
		return PokedexLocation{}, err
	}
	return locationsList, nil
}
