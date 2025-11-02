package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type CategoryRepository interface {
	GetAll() []domain.Category
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll() []domain.Category {
	categories := []domain.Category{}

	r.db.Select(&categories, "SELECT * FROM category ORDER BY id ASC;")

	return categories
}
