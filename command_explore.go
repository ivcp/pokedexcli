package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please input location name")
	}
	data, err := c.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", args[0])

	for _, i := range data.PokemonEncounters {
		fmt.Printf("- %s\n", i.Pokemon.Name)
	}
	return nil
}
