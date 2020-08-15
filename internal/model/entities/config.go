package entities

import "time"

type MockerConfig struct {
	API []*APIConfig
}

type APIConfig struct {
	Routes   []string
	Methods  []string
	Response *APIResponse
}

type APIResponse struct {
	Status  int
	Headers map[string]string
	Body    interface{}
	Time    time.Duration
}
