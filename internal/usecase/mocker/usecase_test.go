package mocker

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/maykonlf/mocker/internal/model/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockHTTPRouter(ctrl)

	t.Run("should register all routes and methods properly", func(t *testing.T) {
		useCase := NewUseCase(m)
		config := &entities.MockerConfig{
			API: []*entities.APIConfig{
				{
					Routes:   []string{"/route-1", "/route-2/2"},
					Methods:  []string{"get", "post", "put"},
					Response: &entities.APIResponse{},
				},
				{
					Routes:   []string{"/routes/2", "/routes/3"},
					Methods:  []string{"get"},
					Response: &entities.APIResponse{},
				},
			},
		}

		m.EXPECT().
			Set(gomock.Eq("/route-1"), gomock.Eq("get"), gomock.Any()).
			Return(nil).
			Times(1)
		m.EXPECT().
			Set(gomock.Eq("/route-1"), gomock.Eq("post"), gomock.Any()).
			Return(nil).
			Times(1)
		m.EXPECT().
			Set(gomock.Eq("/route-1"), gomock.Eq("put"), gomock.Any()).
			Return(nil).
			Times(1)

		m.EXPECT().
			Set(gomock.Eq("/route-2/2"), gomock.Eq("get"), gomock.Any()).
			Return(nil).
			Times(1)
		m.EXPECT().
			Set(gomock.Eq("/route-2/2"), gomock.Eq("post"), gomock.Any()).
			Return(nil).
			Times(1)
		m.EXPECT().
			Set(gomock.Eq("/route-2/2"), gomock.Eq("put"), gomock.Any()).
			Return(nil).
			Times(1)

		m.EXPECT().
			Set(gomock.Eq("/routes/2"), gomock.Eq("get"), gomock.Any()).
			Return(nil).
			Times(1)
		m.EXPECT().
			Set(gomock.Eq("/routes/3"), gomock.Eq("get"), gomock.Any()).
			Return(nil).
			Times(1)

		err := useCase.ConfigRouter(config)
		assert.Nil(t, err)
	})

	t.Run("should return error when set function returns an error", func(t *testing.T) {
		useCase := NewUseCase(m)
		config := &entities.MockerConfig{
			API: []*entities.APIConfig{
				{
					Routes:   []string{"/route-1", "/route-2/2"},
					Methods:  []string{"get", "post", "put"},
					Response: &entities.APIResponse{},
				},
			},
		}
		mockErr := errors.New("some error")
		m.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockErr).Times(1)
		err := useCase.ConfigRouter(config)
		assert.NotNil(t, err)
		assert.Equal(t, err, mockErr)
	})
}
