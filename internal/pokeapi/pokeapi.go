package pokeapi

import (
	"net/http"
	"time"

	"github.com/ivcp/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

const baseUrl = "https://pokeapi.co/api/v2"

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
