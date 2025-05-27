package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sahilrana7582/user-service/internal/domain"
	"github.com/sahilrana7582/user-service/internal/repository"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (email, name, password, role, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query,
		user.Email,
		user.Name,
		user.Password,
		user.Role,
		user.Status,
	).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

func (r *userRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) Update(ctx context.Context, id string, updates *domain.UpdateUserRequest) (*domain.User, error) {
	return nil, nil
}
func (r *userRepo) Delete(ctx context.Context, id string) error {
	return nil
}
func (r *userRepo) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	return nil, nil
}
