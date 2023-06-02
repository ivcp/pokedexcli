package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config, args ...string) error {
	res, err := c.pokeapiClient.ListLocationAreas(c.Next)
	if err != nil {
		return err
	}
	c.Next = res.Next
	c.Prevoius = res.Previous
	for _, i := range res.Results {
		fmt.Println(i.Name)
	}

	return nil
}

func commandMapb(c *config, args ...string) error {
	if c.Prevoius == nil {
		return errors.New("No previous page")
	}
	res, err := c.pokeapiClient.ListLocationAreas(c.Prevoius)
	if err != nil {
		return err
	}
	c.Next = res.Next
	c.Prevoius = res.Previous
	for _, i := range res.Results {
		fmt.Println(i.Name)
	}

	return nil
}
