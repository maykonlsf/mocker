package mocker

import "github.com/maykonlf/mocker/internal/model/entities"

type UseCase interface {
	ConfigRouter(config *entities.MockerConfig) error
}

func NewUseCase(config *entities.MockerConfig) UseCase {
	return &useCase{}
}

type useCase struct {
	httpRouter HTTPRouter
}

func (u *useCase) ConfigRouter(config *entities.MockerConfig) error {
	for _, v := range config.API {
		if err := u.registerConfig(v); err != nil {
			return err
		}
	}

	return nil
}

func (u *useCase) registerConfig(apiConfig *entities.APIConfig) error {
	for i := range apiConfig.Routes {
		for j := range apiConfig.Methods {
			if err := u.httpRouter.Set(apiConfig.Routes[i], apiConfig.Methods[j], apiConfig.Response); err != nil {
				return err
			}
		}
	}

	return nil
}
