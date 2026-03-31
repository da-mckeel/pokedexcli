package main

import (
	"fmt"
	"time"

	"github.com/da-mckeel/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	client := pokeapi.NewClient(time.Second * 10)
	locations, err := client.ListLocations(cfg.Next)
	if err != nil {
		return err
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	return nil
}
