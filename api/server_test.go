package api_test

import (
	"context"
	"go_api_template/api"
	"go_api_template/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewServer(t *testing.T) {
	mockLogger := new(MockLogger)
	mockConfig := utils.Config{Port: "8080"}
	mockDb := new(MockPgPool)

	mockLogger.On("Success", mock.Anything)
	mockLogger.On("Successf", mock.Anything, mock.Anything)
	mockLogger.On("Errorf", mock.Anything, mock.Anything)

	handler := api.NewServer(mockLogger, mockConfig, mockDb)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestRun(t *testing.T) {
	mockLogger := new(MockLogger)
	mockConfig := &utils.Config{Port: "8080"}
	mockDb := new(MockPgPool)

	mockLogger.On("Success", mock.Anything)
	mockLogger.On("Successf", mock.Anything, mock.Anything)
	mockLogger.On("Errorf", mock.Anything, mock.Anything)
	mockLogger.On("Warn", "Server is shutting down...").Return()
	mockDb.On("Close").Return()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := api.Run(ctx, mockLogger, mockConfig, mockDb)
		assert.Nil(t, err)
	}()

	time.Sleep(time.Second)
	cancel()
}

func TestRunFail(t *testing.T) {
	mockLogger := new(MockLogger)
	mockConfig := &utils.Config{Port: "8080"}

	mockLogger.On("Errorf", mock.Anything, mock.Anything)

	err := api.Run(context.Background(), mockLogger, mockConfig, nil)
	assert.NotNil(t, err)
}
