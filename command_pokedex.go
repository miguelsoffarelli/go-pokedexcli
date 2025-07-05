package main

import "fmt"

func commandPokedex(cfg *Config, args ...string) error {
	pokedex := cfg.pokeapiClient.Pokedex
	if len(pokedex) < 1 {
		fmt.Println("You have not caught any Pokemon yet. Use command 'catch <pokemon name>' to catch your first Pokemon.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range pokedex {
		fmt.Printf("- %s\n", name)
	}

	return nil
}
