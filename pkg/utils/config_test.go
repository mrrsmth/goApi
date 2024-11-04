package utils_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go_api_template/pkg/utils"
)

func TestNewConfig(t *testing.T) {
	config := utils.NewConfig(".env.test")

	if config != nil {
		assert.Equal(t, "8080", config.Port)
		assert.Equal(t, "localhost", config.DbHost)
		assert.Equal(t, 5432, config.DbPort)
		assert.Equal(t, "user", config.DbUser)
		assert.Equal(t, "password", config.DbPassword)
		assert.Equal(t, "database", config.DbName)
	} else {
		t.Fail()
	}
}

func TestNewConfigNoFile(t *testing.T) {
	os.Setenv("PORT", "")
	os.Setenv("DB_HOST", "")
	os.Setenv("DB_PORT", "")
	os.Setenv("DB_USER", "")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_NAME", "")

	config := utils.NewConfig("test")

	assert.Nil(t, config)
}

func TestNewConfigInvalidPort(t *testing.T) {
	config := utils.NewConfig(".env.fail")

	assert.Nil(t, config)
}
