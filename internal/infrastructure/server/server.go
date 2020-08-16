package server

import (
	"github.com/maykonlf/mocker/internal/infrastructure/router"
	"github.com/maykonlf/mocker/internal/model/entities"
	"github.com/maykonlf/mocker/internal/usecase/mocker"
)

type Server interface {
	Listen()
}

func NewServer(httpRouter router.Interface, useCase mocker.UseCase, mockerConfig *entities.MockerConfig) Server {
	err := useCase.ConfigRouter(mockerConfig)
	if err != nil {
		panic(err)
	}

	return &server{
		router:  httpRouter,
		useCase: useCase,
	}
}

type server struct {
	router  router.Interface
	useCase mocker.UseCase
}

func (s *server) Listen() {
	s.router.Listen()
}
