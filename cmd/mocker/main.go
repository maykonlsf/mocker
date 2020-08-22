package main

import (
	"net"

	"github.com/maykonlf/mocker/internal/config"
)

func main() {
	container := config.NewContainer()
	listener, err := net.Listen("tcp4", container.GetAddr())
	if err != nil {
		panic(err)
	}

	panic(container.GetServer().Listen(listener))
}
