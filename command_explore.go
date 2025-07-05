package main

import (
	"fmt"

	pokeapi "github.com/miguelsoffarelli/go-pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *Config, area ...string) error {
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
