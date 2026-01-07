package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/middlewares"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewSubcategoryRoutes(mux *http.ServeMux, s services.SubcategoryService) {
	getSubcategoriesByCategory(mux, s)
}

func getSubcategoriesByCategory(mux *http.ServeMux, s services.SubcategoryService) {
	mux.Handle("/subcategory/get-by-category", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			rawUserId := r.Context().Value(middlewares.UserID)

			var userId int

			if rawUserId != nil {
				userId = rawUserId.(int)
			}

			rawId := r.URL.Query().Get("id")

			if rawId == "" {
				w.WriteHeader(400)
				w.Write([]byte("Bad Request"))
				return
			}

			parsedId, _ := strconv.Atoi(rawId)

			categories, err := s.GetSubcategoriesByCategory(parsedId, userId)

			response := dto.GetSubcategoriesByCategoryRes{}

			if err != nil {
				w.WriteHeader(400)
				response.Message = err.Error()
				response.Subcategories = []domain.Subcategory{}
			} else {
				w.WriteHeader(200)
				response.Message = "ok"
				response.Subcategories = categories
			}

			json.NewEncoder(w).Encode(response)

		}
	})))
}
