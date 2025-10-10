package repositories

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type UserRepository interface {
	Create(user domain.User) error
	GetByEmail(email string) domain.User
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByEmail(email string) domain.User {
	usr := domain.User{}

	err := r.db.Get(&usr, "SELECT * FROM user WHERE email = ?", email)

	if err == sql.ErrNoRows {
		return usr
	}

	return usr
}

func (r *userRepository) Create(user domain.User) error {
	fmt.Printf("Inserting user: %+v\n", user)
	_, err := r.db.NamedExec(`
		INSERT INTO user (name, email, password, verification_token, verification_token_expires_at, verified, deleted) 
		VALUES 
		(:name, :email, :password, :verification_token, :verification_token_expires_at, :verified, :deleted)`, user)

	return err
}
