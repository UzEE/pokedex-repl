package api

import (
	"net/http"
	"time"

	"github.com/UzEE/pokedexcli/internal/cache"
)

type Client struct {
	client http.Client
	cache  cache.Cache
}

func NewClient() Client {
	return Client{
		client: http.Client{},
		cache:  cache.NewCache(10 * time.Second),
	}
}
