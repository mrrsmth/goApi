package repositories_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go_api_template/pkg/repositories"
)

type MockDB struct {
	mock.Mock
}

func (db *MockDB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	retArgs := db.Called(ctx, sql, args)
	return retArgs.Get(0).(pgx.Row)
}

func (db *MockDB) Close() {}

type MockRow struct {
	mock.Mock
}

func (mr *MockRow) Scan(dest ...interface{}) error {
	args := mr.Called(dest...)
	return args.Error(0)
}

func TestNewUserRepo(t *testing.T) {
	mockDB := new(MockDB)
	repo := repositories.NewUserRepo(mockDB)
	assert.NotNil(t, repo)
}

func TestGetByID(t *testing.T) {
	mockDB := new(MockDB)
	mockRow := new(MockRow)

	mockRow.On(
		"Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil).Run(func(args mock.Arguments) {
		*(args.Get(0).(*int)) = 1
		*(args.Get(1).(*string)) = "John"
		*(args.Get(2).(*string)) = "Doe"
		*(args.Get(3).(*string)) = "john.doe@example.com"
		*(args.Get(4).(*time.Time)) = time.Now()
		*(args.Get(5).(*time.Time)) = time.Now()
	})

	mockDB.On(
		"QueryRow",
		context.Background(),
		`
		SELECT id, first_name, last_name, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`,
		mock.Anything,
	).Return(mockRow)

	repo := repositories.NewUserRepo(mockDB)
	user, err := repo.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "john.doe@example.com", user.Email)
}

func TestGetByID_Error(t *testing.T) {
	mockDB := new(MockDB)
	mockRow := new(MockRow)

	mockDB.On("QueryRow", context.Background(), mock.Anything, mock.Anything).
		Return(mockRow)
	mockRow.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("database error"))

	repo := repositories.NewUserRepo(mockDB)
	user, err := repo.GetByID(1)

	assert.Error(t, err)
	assert.Nil(t, user)
}
