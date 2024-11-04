package services_test

import (
	"go_api_template/pkg/entities"
	"go_api_template/pkg/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetByID(id int) (*entities.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepo)
	user := &entities.User{
		ID:        1,
		FirstName: "John Doe",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo.On("GetByID", 1).Return(user, nil)

	service := services.NewUserService(mockRepo)
	resultUser, err := service.GetUserByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, resultUser)
	assert.Equal(t, user, resultUser)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByIDError(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockRepo.On("GetByID", 1).Return(nil, assert.AnError)

	service := services.NewUserService(mockRepo)
	resultUser, err := service.GetUserByID(1)

	assert.Error(t, err)
	assert.Nil(t, resultUser)
	mockRepo.AssertExpectations(t)
}
