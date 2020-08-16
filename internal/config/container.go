package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/maykonlf/mocker/internal/infrastructure/router"
	"github.com/maykonlf/mocker/internal/infrastructure/server"
	"github.com/maykonlf/mocker/internal/model/entities"
	"github.com/maykonlf/mocker/internal/usecase/mocker"
	"gopkg.in/yaml.v2"
)

type Container interface {
	GetAddr() string
	GetRouter() router.Interface
	GetMockerUseCase() mocker.UseCase
	GetMockerConfig() *entities.MockerConfig
	GetServer() server.Server
}

func NewContainer() Container {
	flag.Parse()

	config, err := parseConfigFile(*file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &container{
		config: config,
	}
}

type container struct {
	config        *entities.MockerConfig
	router        router.Interface
	mockerUseCase mocker.UseCase
	server        server.Server
}

func (c *container) GetAddr() string {
	return *addr
}

func (c *container) GetRouter() router.Interface {
	if c.router == nil {
		c.router = router.NewRouter(c.GetAddr())
	}

	return c.router
}

func (c *container) GetMockerUseCase() mocker.UseCase {
	if c.mockerUseCase == nil {
		c.mockerUseCase = mocker.NewUseCase(c.GetRouter())
	}

	return c.mockerUseCase
}

func (c *container) GetMockerConfig() *entities.MockerConfig {
	return c.config
}

func (c *container) GetServer() server.Server {
	if c.server == nil {
		c.server = server.NewServer(c.GetRouter(), c.GetMockerUseCase(), c.GetMockerConfig())
	}

	return c.server
}

func parseConfigFile(filename string) (*entities.MockerConfig, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var rawConfig router.MockerConfig
	err = yaml.Unmarshal(fileData, &rawConfig)
	if err != nil {
		return nil, err
	}

	return rawConfig.ToEntityModel()
}
