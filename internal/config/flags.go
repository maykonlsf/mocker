package config

import "flag"

var (
	file = flag.String("f", "mocker.yaml", "mocker.yaml config file path")
	addr = flag.String("addr", ":8081", "server API address")
)
