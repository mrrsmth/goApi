package domain_test

import (
	"context"
	"testing"

	"go_api_template/domain"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Success(msg string) {
	m.Called(msg)
}

func (m *MockLogger) Successf(msg string, args ...interface{}) {
	m.Called(msg, args)
}

func (m *MockLogger) Info(msg string) {
	m.Called(msg)
}

func (m *MockLogger) Infof(msg string, args ...interface{}) {
	m.Called(msg, args)
}

func (m *MockLogger) Error(msg string) {
	m.Called(msg)
}

func (m *MockLogger) Errorf(msg string, args ...interface{}) {
	m.Called(msg, args)
}

func (m *MockLogger) Warn(msg string) {
	m.Called(msg)
}

func (m *MockLogger) Warnf(msg string, args ...interface{}) {
	m.Called(msg, args)
}

type MockConfig struct {
	mock.Mock
}

type MockPgPool struct {
	mock.Mock
}

func (m *MockPgPool) Close() {
	m.Called()
}

func (db *MockPgPool) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	retArgs := db.Called(ctx, sql, args)
	return retArgs.Get(0).(pgx.Row)
}

func TestUserDomain(t *testing.T) {
	logger := new(MockLogger)
	logger.On("Success", "Bootstrapped user domain.")

	config := new(MockConfig)
	db := new(MockPgPool)

	userHandler := domain.UserDomain(logger, config, db)

	assert.NotNil(t, userHandler)
	logger.AssertExpectations(t)
}
