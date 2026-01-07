package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type SubcategoryRepository interface {
	GetSubcategoriesByCategory(categoryId int) ([]domain.Subcategory, error)
}

type subcategoryRepository struct {
	db *sqlx.DB
}

func NewSubcategoryRepository(db *sqlx.DB) SubcategoryRepository {
	return &subcategoryRepository{db: db}
}

func (r *subcategoryRepository) GetSubcategoriesByCategory(categoryId int) ([]domain.Subcategory, error) {
	var ret []domain.Subcategory

	err := r.db.Select(&ret, "SELECT * FROM subcategory WHERE category_id = ?", categoryId)

	return ret, err

}
