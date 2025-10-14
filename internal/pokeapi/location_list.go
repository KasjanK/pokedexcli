package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

)

func (c *Client) FetchLocationsList(pageURL *string) (PokedexLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	
	if val, ok := c.cache.Get(url); ok {
		locationsList := PokedexLocation{}
		err := json.Unmarshal(val, &locationsList)
		if err != nil {
			return PokedexLocation{}, err
		}
		return locationsList, err
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokedexLocation{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokedexLocation{}, err
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

	c.cache.Add(url, body)
	return locationsList, nil
}
