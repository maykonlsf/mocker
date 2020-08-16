package router

import "github.com/maykonlf/mocker/internal/model/entities"

type Interface interface {
	Set(route, method string, response *entities.APIResponse) error
	Listen()
}
