package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocation(pageURL string) (Location, error) {
	url := baseURL + "/location-area/"
	if pageURL != "" {
		url += pageURL
	}
	
	if val, ok := c.cache.Get(url); ok {
		pokemonList := Location{}
		err := json.Unmarshal(val, &pokemonList)
		if err != nil {
			return Location{}, err
		}
		return pokemonList, err
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("Error reading response body: %w", err)
	}

	location := Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, body)
	return location, nil
}
