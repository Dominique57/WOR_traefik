package wor_traefik

import (
	"context"
	"net/http"
)

type Config struct {
}

func CreateConfig() *Config {
	return &Config{}
}

type WorDemo struct {
	next http.Handler
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &WorDemo{
		next: next,
		name: name,
	}, nil
}

func (a *WorDemo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	a.next.ServeHTTP(rw, req)
}
