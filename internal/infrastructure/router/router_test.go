package router

import (
	"testing"

	"github.com/maykonlf/mocker/internal/model/entities"
	"github.com/stretchr/testify/assert"
)

func TestSetRoute(t *testing.T) {
	r := NewRouter(":8081")
	t.Run("should accept add new route", func(t *testing.T) {
		err := r.Set("/route", "get", &entities.APIResponse{
			Status:  200,
			Body:    "",
			Headers: map[string]string{},
			Time:    0,
		})

		assert.Nil(t, err)

		err = r.Set("/route", "get", &entities.APIResponse{
			Status:  200,
			Body:    "",
			Headers: map[string]string{},
			Time:    0,
		})

		assert.NotNil(t, err)
	})
}
