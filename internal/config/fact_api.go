package config

import (
	"fmt"
	"os"
)

const (
	APIKeyEnvName = "API_KEY"
)

// FactAPI ...
type FactAPI interface {
	GetAPIKey() string
}

type factAPI struct {
	APIKey string
}

// NewFactAPI ...
func NewFactAPI() (FactAPI, error) {
	APIKey := os.Getenv(APIKeyEnvName)
	if len(APIKey) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", APIKeyEnvName)
	}

	return &factAPI{
		APIKey: APIKey,
	}, nil
}

// GetAPIKey ...
func (fa *factAPI) GetAPIKey() string {
	return fa.APIKey
}
