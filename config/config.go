// Package config is responsible for parsing config file into Config.
package config

type Config struct {
	Port      int
	Publisher Publisher
}

type Publisher struct {
	Name string
}
