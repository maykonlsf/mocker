package main

import (
	"github.com/maykonlf/mocker/internal/config"
)

func main() {
	container := config.NewContainer()
	panic(container.GetServer().Listen())
}
