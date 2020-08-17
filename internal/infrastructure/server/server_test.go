package server

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/maykonlf/mocker/internal/model/entities"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp/fasthttputil"
	"testing"
)

func TestNewServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	routerMock := NewMockRouter(ctrl)
	useCaseMock := NewMockUseCase(ctrl)

	config := &entities.MockerConfig{}

	t.Run("should return error when fails to create new server", func(t *testing.T) {
		configError := errors.New("config error: invalid config file")
		useCaseMock.EXPECT().ConfigRouter(gomock.Eq(config)).Return(configError).Times(1)
		server, err := NewServer(routerMock, useCaseMock, config)
		assert.Nil(t, server)
		assert.Equal(t, configError, err)
	})

	t.Run("should return the server instance and no error when can create server whitout error", func(t *testing.T) {
		useCaseMock.EXPECT().ConfigRouter(config).Return(nil).Times(1)
		server, err := NewServer(routerMock, useCaseMock, config)
		assert.NotNil(t, server)
		assert.Nil(t, err)
	})
}

func TestListen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	routerMock := NewMockRouter(ctrl)
	useCaseMock := NewMockUseCase(ctrl)

	config := &entities.MockerConfig{}

	t.Run("should call route listener", func(t *testing.T) {
		useCaseMock.EXPECT().ConfigRouter(config).Return(nil).Times(1)
		server, err := NewServer(routerMock, useCaseMock, config)
		assert.NotNil(t, server)
		assert.Nil(t, err)

		routerMock.EXPECT().Listen(gomock.Any()).Return(nil)
		err = server.Listen(fasthttputil.NewInmemoryListener())
		assert.Nil(t, err)
	})
}
