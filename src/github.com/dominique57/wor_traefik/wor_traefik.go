package wor_traefik

import (
	"context"
	"fmt"
	"net/http"
)

type Config struct {
    Foo string
}

func CreateConfig() *Config {
	return &Config{
        Foo: "",
	}
}

type WorDemo struct {
	next http.Handler
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
    if len(config.Foo) == 0 {
        return nil, fmt.Errorf("Config Foo cannot be empty !");
    }

	return &WorDemo{
		next: next,
		name: name,
	}, nil
}

func (a *WorDemo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	a.next.ServeHTTP(rw, req)
}
