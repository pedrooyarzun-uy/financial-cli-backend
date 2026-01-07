package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/middlewares"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewSubcategoryRoutes(mux *http.ServeMux, s services.SubcategoryService) {
	getSubcategoriesByCategory(mux, s)
}

func getSubcategoriesByCategory(mux *http.ServeMux, s services.SubcategoryService) {
	mux.Handle("/subcategory/get-by-category", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			rawId := r.URL.Query().Get("id")
			parsedId, _ := strconv.Atoi(rawId)

			categories, err := s.GetSubcategoriesByCategory(parsedId)

			response := dto.GetSubcategoriesByCategoryRes{}

			if err != nil {
				w.WriteHeader(400)
				response.Message = err.Error()
			} else {
				w.WriteHeader(200)
				response.Message = "ok"
			}

			response.Subcategories = categories
			json.NewEncoder(w).Encode(response)

		}
	})))
}
