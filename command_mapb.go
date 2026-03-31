package main

import (
	"fmt"
	"time"

	"github.com/da-mckeel/pokedexcli/internal/pokeapi"
)

func commandMapB(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	client := pokeapi.NewClient(time.Second * 10)
	locations, err := client.ListLocations(cfg.Previous)
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
