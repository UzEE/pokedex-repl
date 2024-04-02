package api

import (
	"net/http"
	"time"

	"github.com/UzEE/pokedexcli/internal/cache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	client http.Client
	cache  cache.Cache
}

func NewClient() Client {
	return Client{
		client: http.Client{},
		cache:  cache.NewCache(60 * time.Second),
	}
}
