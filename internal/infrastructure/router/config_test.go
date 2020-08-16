package router

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfigParser(t *testing.T) {
	t.Run("should parse the valid config to model entity", func(t *testing.T) {
		config := &MockerConfig{
			API: []*APIConfig{
				{
					Routes:  []string{"/route-1", "/route-2"},
					Methods: []string{"get", "post"},
					Response: &APIResponse{
						Status:  200,
						Body:    "{}",
						Headers: map[string]string{"Content-Type": "application/json"},
						Time:    "100ms",
					},
				},
			},
		}

		entityConfig, err := config.ToEntityModel()
		assert.NotNil(t, entityConfig)
		assert.Nil(t, err)

		assert.Equal(t, config.API[0].Methods, entityConfig.API[0].Methods)
		assert.Equal(t, config.API[0].Routes, entityConfig.API[0].Routes)
		assert.Equal(t, config.API[0].Response.Status, entityConfig.API[0].Response.Status)
		assert.Equal(t, config.API[0].Response.Body, entityConfig.API[0].Response.Body)
		assert.Equal(t, config.API[0].Response.Headers, entityConfig.API[0].Response.Headers)
		assert.Equal(t, 100*time.Millisecond, entityConfig.API[0].Response.Time)
	})

	t.Run("should return error when tries to parse an invalid config to entity", func(t *testing.T) {
		config := &MockerConfig{
			API: []*APIConfig{
				{
					Routes:  []string{"/route-1", "/route-2"},
					Methods: []string{"get", "post"},
					Response: &APIResponse{
						Status:  200,
						Body:    "{}",
						Headers: map[string]string{"Content-Type": "application/json"},
						Time:    "invalid",
					},
				},
			},
		}

		_, err := config.ToEntityModel()
		assert.NotNil(t, err)
	})
}
