package main

import (
	"time"

	"github.com/ivcp/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	repl(&cfg)
}
