package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokeapi "github.com/miguelsoffarelli/go-pokedexcli/internal/pokeapi"
	pokecache "github.com/miguelsoffarelli/go-pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(2 * time.Minute)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scan.Scan()
		input := scan.Text()
		clearInput := cleanInput(input)
		command, ok := supportedCommands[clearInput[0]]
		if ok {
			command.callback(&config)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	words := make([]string, 0)

	for _, word := range strings.Fields(text) {
		words = append(words, strings.ToLower(word))
	}

	return words
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	fmt.Printf(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}

func commandMap(c *Config) error {
	locationData, err := pokeapi.FetchLocationArea(c.next, cache)
	if err != nil {
		return err
	}

	for _, area := range locationData.Results {
		fmt.Println(area.Name)
	}

	c.next, c.previous = locationData.Next, locationData.Previous

	return nil
}

func commandMapb(c *Config) error {
	if c.previous == nil {
		fmt.Println("you're in the first page")
		return nil
	}

	locationData, err := pokeapi.FetchLocationArea(*c.previous, cache)
	if err != nil {
		return err
	}

	for _, area := range locationData.Results {
		fmt.Println(area.Name)
	}

	c.next, c.previous = locationData.Next, locationData.Previous

	return nil
}

type CliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

type Config struct {
	next     string
	previous *string
}

type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var supportedCommands = map[string]CliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Show available commands",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Show 20 map locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Show previous 20 map locations",
		callback:    commandMapb,
	},
}

var config = Config{
	next:     "https://pokeapi.co/api/v2/location-area/",
	previous: nil,
}
