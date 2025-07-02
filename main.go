package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scan.Scan()
		input := scan.Text()
		clearInput := cleanInput(input)
		command, ok := supportedCommands[clearInput[0]]
		if ok {
			command.callback()
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}

func commandMap() error {
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands = map[string]cliCommand{
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
		description: "Show map locations",
		callback:    commandMap,
	},
}
