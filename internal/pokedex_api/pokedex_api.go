package pokedex_api

import (
	"encoding/json"
	"log"
	"net/http"
)

type LocationArea struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocations(url string) (locations LocationArea, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error: sending GET request", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, resp.Body)
	}
	var locationArea LocationArea
	if err := json.NewDecoder(resp.Body).Decode(&locationArea); err != nil {
		log.Fatal("error: decoding respose json", err)
	}
	return locationArea, nil
}

func Locations(locArea LocationArea) (locations []string) {
	locs := make([]string, len(locArea.Results))
	for i, result := range locArea.Results {
		locs[i] = result.Name
	}
	return locs
}
