package repositories

import (
	"context"
	"go_api_template/pkg/entities"
	"go_api_template/pkg/utils"
)

type IUserRepo interface {
	GetByID(id int) (*entities.User, error)
}

type UserRepo struct {
	db utils.PgPool
}

// NewUserRepo creates a new UserRepo. (takes in db connection pool)
func NewUserRepo(db utils.PgPool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetByID(id int) (*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	user := &entities.User{}
	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
