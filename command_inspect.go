package main

import "fmt"

func commandInspect(cfg *Config, pokemon ...string) error {
	if len(pokemon) != 1 {
		return fmt.Errorf("must provide one Pokemon name")
	}

	name := pokemon[0]

	data, ok := cfg.pokeapiClient.Pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf(`
	Name: %s
	Height: %v
	Weight: %v
	Stats:
		-hp: %v
		-attack: %v
		-defense: %v
		-special-attack: %v
		-special-defense: %v
		-speed: %v
	Types:
`,
		data.Name,
		data.Height,
		data.Weight,
		data.Stats[0].BaseStat, // hp
		data.Stats[1].BaseStat, // attack
		data.Stats[2].BaseStat, // defense
		data.Stats[3].BaseStat, // special-attack
		data.Stats[4].BaseStat, // special-defense
		data.Stats[5].BaseStat, // speed
	)

	for i := range data.Types {
		fmt.Printf("		- %v\n", data.Types[i].Type.Name)
	}

	return nil
}
