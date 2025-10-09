package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/db"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type UserRepository interface {
	Create(user domain.User) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user domain.User) error {
	_, err := db.DB.NamedExec(`
		INSERT INTO user (name, email, password, verfication_token, verification_token_expires_at, verified, created_at, deleted) 
		VALUES 
		(:Name, :Email, :Password, :VerificationToken, :VerificationTokenExpiresAt, :Verified, :CreatedAt, :Deleted)`, user)

	return err
}
