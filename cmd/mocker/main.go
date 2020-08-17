package main

import (
	"github.com/maykonlf/mocker/internal/config"
	"net"
)

func main() {
	container := config.NewContainer()
	listener, err := net.Listen("tcp4", container.GetAddr())
	if err != nil {
		panic(err)
	}

	panic(container.GetServer().Listen(listener))
}
