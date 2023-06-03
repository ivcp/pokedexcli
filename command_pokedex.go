package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("invalid command")
	}
	if len(c.caughtPokemon) == 0 {
		return errors.New("Nothing in your pokedex.")
	}
	fmt.Println("Your pokedex:")
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
