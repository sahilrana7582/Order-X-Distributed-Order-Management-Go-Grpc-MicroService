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
	var user domain.User

	query := `
		SELECT id, email, name, password, role, status, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found with id: %s", id)
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) Update(ctx context.Context, id string, updates *domain.UpdateUserRequest) (*domain.User, error) {

	query := `
		UPDATE users
		SET name = COALESCE(NULLIF($1, ''), name),
		email = COALESCE(NULLIF($2, ''), email),
		password = COALESCE(NULLIF($3, ''), password),
		role = COALESCE(NULLIF($4, ''), role),
		status = COALESCE(NULLIF($5, ''), status),
		updated_at = NOW()
		WHERE id = $6
		RETURNING id, email, name, password, role, status, created_at, updated_at
	`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, updates.Name, updates.Email, updates.Password, updates.Role, updates.Status, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found with id: %s", id)
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return &user, nil
}
func (r *userRepo) Delete(ctx context.Context, id string) error {

	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user with id %s: %w", id, err)
	}
	return nil
}
func (r *userRepo) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	query := `
		SELECT id, email, name, password, role, status, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Password,
			&user.Role,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %w", err)
	}

	return users, nil
}
