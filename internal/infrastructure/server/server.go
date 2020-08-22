package server

import (
	"net"

	"github.com/maykonlf/mocker/internal/infrastructure/router"
	"github.com/maykonlf/mocker/internal/model/entities"
	"github.com/maykonlf/mocker/internal/usecase/mocker"
)

type Server interface {
	Listen(ln net.Listener) error
}

func NewServer(httpRouter router.Router, useCase mocker.UseCase, mockerConfig *entities.MockerConfig) (Server, error) {
	err := useCase.ConfigRouter(mockerConfig)
	if err != nil {
		return nil, err
	}

	return &server{
		router:  httpRouter,
		useCase: useCase,
	}, nil
}

type server struct {
	router  router.Router
	useCase mocker.UseCase
}

func (s *server) Listen(ln net.Listener) error {
	return s.router.Listen(ln)
}
