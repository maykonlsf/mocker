package main

import (
	"flag"

	"github.com/maykonlf/mocker/internal/config"
)

func main() {
	flag.Parse()

	container := config.NewContainer()
	container.GetServer().Listen()
}
