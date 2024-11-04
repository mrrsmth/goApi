package api_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"go_api_template/api"
)

func TestAddRoutes(t *testing.T) {
	logger := new(MockLogger)
	config := new(MockConfig)
	db := new(MockPgPool)

	logger.On("Success", "Bootstrapped user domain.").Return()
	logger.On("Success", "Routes added.").Return()

	mux := http.NewServeMux()
	api.AddRoutes(mux, logger, config, db)

	logger.AssertCalled(t, "Success", "Bootstrapped user domain.")
	logger.AssertCalled(t, "Success", "Routes added.")

	assert.Equal(t, mux, mux)
}
