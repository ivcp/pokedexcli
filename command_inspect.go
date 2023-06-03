package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please input valid pokemon name")
	}
	val, isCaught := c.caughtPokemon[args[0]]

	if !isCaught {
		return errors.New("You haven't cought that pokemon.")
	}

	fmt.Printf("Name: %s\nHeight: %d\nStats:\n", val.Name, val.Height)

	for _, stat := range val.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.StatName, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range val.Types {
		fmt.Printf(" -%s\n", t.Type.TypeName)
	}

	return nil
}
