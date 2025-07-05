package main

import (
	"fmt"

	pokeapi "github.com/miguelsoffarelli/go-pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *Config, area ...string) error {
	// area argument must be at least one, but can be more
	// Example: "explore" is not valid. "explore area-name-1" is valid.
	// "explore area-name-1 area-name-2 area-name-n" is also valid.
	if len(area) < 1 {
		return fmt.Errorf("you must provide at least one location name")
	}

	// Loop over area slice in case there's more than one
	// to allow the user to explore multiple areas with one command
	for _, a := range area {
		url := pokeapi.BaseURL + "/location-area/" + a

		areaList, err := cfg.pokeapiClient.ListArea(&url)
		if err != nil {
			return err
		}

		fmt.Printf("Exploring %s...", a)
		fmt.Println("Found Pokemon:")

		for _, p := range areaList.PokemonEncounters {
			fmt.Printf("- %s\n", p.Pokemon.Name)
		}
	}
	return nil
}
