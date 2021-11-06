package config

import (
	"os"
	"strings"
)

// Config contains all information about the service
type Config struct {
	// Credentials is a map that stores username:password rows
	Credentials map[string]string
}

func New(pathToCreds string) (*Config, error) {
	creds, err := os.ReadFile(pathToCreds)
	if err != nil {
		return nil, err
	}
	credsStr := string(creds)
	credsByRow := strings.Split(credsStr, "\n")

	credentials := make(map[string]string, len(credsByRow))

	for _, credsRow := range credsByRow {
		data := strings.Split(credsRow, ":")
		user, password := data[0], data[1]
		credentials[user] = password
	}

	config := new(Config)
	config.Credentials = credentials

	return config, nil
}
