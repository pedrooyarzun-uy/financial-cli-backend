package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type CategoryRepository interface {
	GetAll(userId int) []domain.Category
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll(userId int) []domain.Category {
	categories := []domain.Category{}

	r.db.Select(&categories, "SELECT * FROM category WHERE user_id = ? OR user_id IS NULL ORDER BY name ASC;", userId)

	return categories
}
