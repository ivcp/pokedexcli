package main

import "fmt"

func commandHelp(c *config) error {
	fmt.Print("\nWelcome to the pokedex!\nUsage:\n\n")
	for _, i := range getCommands() {
		fmt.Printf("%s: %s\n\n", i.name, i.description)
	}
	return nil
}
