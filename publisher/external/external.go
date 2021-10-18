package external

type Client interface {
	Get() error
}

func New(name string) Client {
	return &client{
		name: name,
	}
}

type client struct {
	name string
}

func (c client) Get() error {
	return nil
}
