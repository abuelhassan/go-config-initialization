package publisher

import (
	"errors"
	"log"

	"github.com/abuelhassan/go-config-initialization/config"
	"github.com/abuelhassan/go-config-initialization/initializer"
	"github.com/abuelhassan/go-config-initialization/publisher/external"
)

// Don't use directly, for easier unit tests ;)
var client external.Client

type Publisher interface {
	Publish() error
}

func New() Publisher {
	return &publisher{
		client: client,
	}
}

type publisher struct {
	client external.Client
}

func (p publisher) Publish() error {
	return p.client.Get()
}

func initialize(cfg config.Config) error {
	if cfg.Publisher.Name == "" {
		return errors.New("empty publisher name")
	}

	client = external.New(cfg.Publisher.Name)
	log.Println("publisher initialized")

	return nil
}

func init() {
	initializer.Register(initialize)
}
