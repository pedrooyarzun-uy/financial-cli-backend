package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/middlewares"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewCategoryRoutes(mux *http.ServeMux, s services.CategoryService) {
	getAllCategories(mux, s)
}

func getAllCategories(mux *http.ServeMux, s services.CategoryService) {
	mux.Handle("/category/get-all", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			userIdVal := r.Context().Value(middlewares.UserID)

			var userId int

			if userIdVal != nil {
				userId = userIdVal.(int)
			}

			res := s.GetAll(userId)

			response := dto.GetAllCategoryRes{}

			w.WriteHeader(200)
			response.Message = "ok"
			response.Categories = res
			json.NewEncoder(w).Encode(response)
		}
	})))
}
