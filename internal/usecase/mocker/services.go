package mocker

import "github.com/maykonlf/mocker/internal/model/entities"

type HTTPRouter interface {
	Set(route, method string, response *entities.APIResponse) error
}
