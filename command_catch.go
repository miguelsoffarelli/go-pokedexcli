package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func commandCatch(cfg *Config, pokemon ...string) error {
	// In this case (unlike in commandExplore), pokemon argument must be one
	// Example: "catch" is not valid, "catch pokemon-name" is valid,
	// "catch pokemon-name-1 pokemon-name-2 pokemon-name-n" is NOT valid in this case
	if len(pokemon) != 1 {
		return fmt.Errorf("you must provide one pokemon name")
	}

	name := pokemon[0]

	pokemonList, err := cfg.pokeapiClient.ListPokemon(&name)
	if err != nil {
		return err
	}

	baseXP := float64(pokemonList.BaseExperience)
	difficulty := getDifficulty(cfg)
	chance := 1.0 / (baseXP / difficulty)
	// Avoid chance being lower than 0.1 or higher than 0.99
	chance = math.Min(math.Max(chance, 0.1), 0.99)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if rand.Float64() <= chance {
		cfg.pokeapiClient.Pokedex[name] = pokemonList
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func getDifficulty(cfg *Config) float64 {
	switch cfg.difficulty {
	case "easy":
		return 50.0
	case "medium":
		return 40.0
	case "hard":
		return 30.0
	default:
		return 50.0
	}
}
