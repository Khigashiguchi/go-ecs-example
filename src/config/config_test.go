package config_test

import (
	"testing"

	"os"

	"github.com/Khigashiguchi/go-ecs-example/src/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		name     string
		expected config.Config
	}{
		{
			name: "localhost",
			expected: config.Config{
				DB: config.DBConfig{
					User:     "sample_user",
					Password: "sample_password",
					Host:     "db.locahost",
					Port:     3306,
					Name:     "sample",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			os.Setenv("DB_USER", "sample_user")
			os.Setenv("DB_PASSWORD", "sample_password")
			os.Setenv("DB_HOST", "db.locahost")
			os.Setenv("DB_PORT", "3306")
			os.Setenv("DB_NAME", "sample")

			res, err := config.NewConfig()

			assert.Equal(t, nil, err)
			assert.Equal(t, c.expected.DB.User, res.DB.User)
			assert.Equal(t, c.expected.DB.Password, res.DB.Password)
			assert.Equal(t, c.expected.DB.Host, res.DB.Host)
			assert.Equal(t, c.expected.DB.Port, res.DB.Port)
			assert.Equal(t, c.expected.DB.Name, res.DB.Name)
		})
	}
}
