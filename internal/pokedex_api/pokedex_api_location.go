package pokedex_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/7minutech/pokedex/internal/pokecache"
)

type LocationAreaPage struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var cache = pokecache.NewCache(time.Second * 5)

func GetLocations(url string) (locations LocationAreaPage, err error) {
	var locationArea LocationAreaPage
	if value, found := cache.Get(url); found {
		if err := json.Unmarshal(value, &locationArea); err != nil {
			return LocationAreaPage{}, fmt.Errorf("error ummarshaling cache data -> %w", err)
		}
		return locationArea, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaPage{}, fmt.Errorf("error: sending GET request -> %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		b, _ := io.ReadAll(resp.Body)
		return LocationAreaPage{}, fmt.Errorf("error: status code: %d: %s", resp.StatusCode, string(b))
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaPage{}, fmt.Errorf("error: reading response body -> %w", err)
	}
	cache.Add(url, data)
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationAreaPage{}, fmt.Errorf("error: unmarshaling data -> %w", err)
	}
	return locationArea, nil

}

func Locations(locArea LocationAreaPage) (locations []string) {
	locs := make([]string, len(locArea.Results))
	for i, result := range locArea.Results {
		locs[i] = result.Name
	}
	return locs
}
