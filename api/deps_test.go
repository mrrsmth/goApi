package api_test

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"

	"go_api_template/pkg/utils"
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
	utils.Config
}

type MockPgPool struct {
	mock.Mock
}

func (m *MockPgPool) Close() {
	m.Called()
}

func (m *MockPgPool) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	m.Called(sql, args)
	return nil
}
