package initializer

import "github.com/abuelhassan/go-config-initialization/config"

var (
	inits []func(config.Config) error
)

func Register(init func(config.Config) error) {
	inits = append(inits, init)
}

type Initializer interface {
	InitializePackages(cfg config.Config) error
}

func New() Initializer {
	return &initializer{}
}

type initializer struct{}

func (initializer) InitializePackages(cfg config.Config) error {
	for _, f := range inits {
		err := f(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}
