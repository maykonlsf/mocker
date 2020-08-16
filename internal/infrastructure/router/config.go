package router

import (
	"fmt"
	"time"

	"github.com/maykonlf/mocker/internal/model/entities"
)

type MockerConfig struct {
	API []*APIConfig `yaml:"api"`
}

func (c *MockerConfig) ToEntityModel() (*entities.MockerConfig, error) {
	parsedConfig := &entities.MockerConfig{
		API: make([]*entities.APIConfig, len(c.API)),
	}

	for i := range c.API {
		config, err := c.API[i].ToEntity()
		if err != nil {
			return nil, err
		}

		parsedConfig.API[i] = config
	}

	return parsedConfig, nil
}

type APIConfig struct {
	Routes   []string     `yaml:"routes"`
	Methods  []string     `yaml:"methods"`
	Response *APIResponse `yaml:"response"`
}

func (a *APIConfig) ToEntity() (*entities.APIConfig, error) {
	response, err := a.Response.ToEntity()
	if err != nil {
		return nil, err
	}

	return &entities.APIConfig{
		Routes:   a.Routes,
		Methods:  a.Methods,
		Response: response,
	}, nil
}

type APIResponse struct {
	Status  int               `yaml:"status"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
	Time    string            `yaml:"time"`
}

func (r *APIResponse) ToEntity() (*entities.APIResponse, error) {
	duration, err := time.ParseDuration(r.Time)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response duration: %s", err)
	}

	return &entities.APIResponse{
		Body:    r.Body,
		Headers: r.Headers,
		Status:  r.Status,
		Time:    duration,
	}, nil
}
