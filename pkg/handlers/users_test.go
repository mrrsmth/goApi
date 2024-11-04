package handlers_test

import (
	"errors"
	"go_api_template/pkg/entities"
	"go_api_template/pkg/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserByID(id int) (*entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func TestGetUserByID(t *testing.T) {
	mockUserService := new(MockUserService)
	mockLogger := new(MockLogger)
	userHandler := handlers.NewUserHandler(mockUserService)

	// Set up the mock UserService to return a specific user for ID 1
	mockUserService.On("GetUserByID", 1).Return(&entities.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}, nil)

	// Set up the mock Logger to expect an Errorf call
	mockLogger.On("Errorf", mock.Anything, mock.Anything)
	// Set up the mock Logger to expect an Error call
	mockLogger.On("Error", "Invalid ID")

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/{id}", userHandler.GetUserByID(mockLogger))

	req := httptest.NewRequest("GET", "/api/users/1", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	mockUserService.AssertCalled(t, "GetUserByID", 1)
	mockLogger.AssertNotCalled(t, "Errorf")
	mockLogger.AssertNotCalled(t, "Error")

	expected := `{"status":200,"data":{"id":1,"firstName":"John","lastName":"Doe","email":"john.doe@example.com","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}}` + "\n"

	assert.Equal(t, expected, rr.Body.String())
}

func TestGetUserByIDInvalidID(t *testing.T) {
	mockUserService := new(MockUserService)
	mockLogger := new(MockLogger)
	userHandler := handlers.NewUserHandler(mockUserService)

	// Set up the mock Logger to expect an Errorf call
	mockLogger.On("Errorf", mock.Anything, mock.Anything)
	// Set up the mock Logger to expect an Error call
	mockLogger.On("Error", "Invalid ID")

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/{id}", userHandler.GetUserByID(mockLogger))

	req := httptest.NewRequest("GET", "/api/users/invalid", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	mockUserService.AssertNotCalled(t, "GetUserByID")
	mockLogger.AssertNotCalled(t, "Errorf")
	mockLogger.AssertCalled(t, "Error", "Invalid ID")
}

func TestGetUserByIDError(t *testing.T) {
	mockUserService := new(MockUserService)
	mockLogger := new(MockLogger)
	userHandler := handlers.NewUserHandler(mockUserService)

	// Set up the mock UserService to return an error
	mockUserService.On("GetUserByID", 1).Return((*entities.User)(nil), errors.New("Something went wrong"))
	// Set up the mock Logger to expect an Errorf call
	mockLogger.On("Errorf", mock.Anything, mock.Anything)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/{id}", userHandler.GetUserByID(mockLogger))

	req := httptest.NewRequest("GET", "/api/users/1", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	mockUserService.AssertCalled(t, "GetUserByID", 1)
	mockLogger.AssertCalled(t, "Errorf", "Something went wrong getting user by ID: %v", mock.Anything)
}
