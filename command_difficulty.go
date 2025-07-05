package main

import "fmt"

func commandDifficulty(cfg *Config, dif ...string) error {
	if len(dif) != 1 {
		return fmt.Errorf("must select one difficulty level. Options: easy - medium - hard")
	}

	difficulty := cleanInput(dif[0])
	cfg.difficulty = difficulty[0]

	return nil
}
