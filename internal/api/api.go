package api

import (
	"crypto/rand"
	"crypto/rsa"

	clientsdb "github.com/pietervdwerk/tasksapi/internal/clients"
	tasksdb "github.com/pietervdwerk/tasksapi/internal/tasks"
)

var _ StrictServerInterface = (*API)(nil)

type API struct {
	conf        *APIConfig
	tasksRepo   tasksdb.Querier
	clientsRepo clientsdb.Querier
}

type APIConfigFunc func(*APIConfig)

type APIConfig struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewDefaultAPIConfig creates a new APIConfig with a default RSA key pair
func NewDefaultAPIConfig() (*APIConfig, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &APIConfig{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}, nil
}

// NewAPI creates a new API instance
func NewAPI(tasksRepo tasksdb.Querier, clientsRepo clientsdb.Querier, opts ...APIConfigFunc) (*API, error) {
	defaultConf, err := NewDefaultAPIConfig()
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(defaultConf)
	}

	return &API{
		conf:        defaultConf,
		tasksRepo:   tasksRepo,
		clientsRepo: clientsRepo,
	}, nil
}

func WithPrivateKey(privateKey *rsa.PrivateKey) APIConfigFunc {
	return func(c *APIConfig) {
		c.privateKey = privateKey
	}
}
