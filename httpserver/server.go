package httpserver

import (
	"net/http"

	"go.uber.org/zap"
)

// Service is a struct representing service.
// It contains a http.Server and zap.Logger that is common for all project
type Service struct {
	Server *http.Server
	Logger zap.Logger
}

// New creates and returns new object of Service
func New() (*Service, error) {
	server := http.Server{
		Addr:    ":9000",
		Handler: nil,
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	service := new(Service)
	service.Server = &server
	service.Logger = *logger

	return service, nil
}
