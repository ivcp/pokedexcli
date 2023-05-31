package main

import (
	"github.com/ivcp/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Prevoius      *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	repl(&cfg)
}
