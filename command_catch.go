package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please input pokemon name")
	}
	data, err := c.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		msg := fmt.Sprintf("no pokemon named %s", args[0])
		return errors.New(msg)
	}

	if pokemon, exists := c.caughtPokemon[data.Name]; exists {
		msg := fmt.Sprintf("%s already caught!", pokemon.Name)
		return errors.New(msg)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", data.Name)

	chance := rand.Intn(500)
	fmt.Println(chance)
	if chance < data.BaseExperience {
		fmt.Printf("%s escaped!\n", data.Name)
	} else {
		fmt.Printf("%s was caught!\n", data.Name)
		c.caughtPokemon[data.Name] = data
	}

	return nil
}
