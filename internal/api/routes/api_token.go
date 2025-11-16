package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/middlewares"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewApiToken(mux *http.ServeMux, s services.ApiTokenService) {
	createApiToken(mux, s)
	getAllApiTokens(mux, s)
}

func createApiToken(mux *http.ServeMux, s services.ApiTokenService) {
	mux.Handle("/api-token/create", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			userIdVal := r.Context().Value(middlewares.UserID)

			var userId int

			if userIdVal != nil {
				userId = userIdVal.(int)
			}

			var body dto.CreateApiTokenReq

			err := json.NewDecoder(r.Body).Decode(&body)

			if err != nil {
				w.WriteHeader(402)
				w.Write([]byte("Something went wrong"))
				return
			}

			if body.Keyword == "" || body.Name == "" {
				w.WriteHeader(400)
				w.Write([]byte("Bad request"))
				return
			}

			response := dto.GenericRes{}

			err = s.Create(userId, body)

			if err == nil {
				w.WriteHeader(200)
				response.Message = "ok"
				json.NewEncoder(w).Encode(response)
				return
			}

			w.WriteHeader(400)
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
		}
	})))
}

func getAllApiTokens(mux *http.ServeMux, s services.ApiTokenService) {
	mux.Handle("/api-token/get-all", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			userIdVal := r.Context().Value(middlewares.UserID)

			var userId int

			if userIdVal != nil {
				userId = userIdVal.(int)
			}

			ret := s.GetAll(userId)

			response := dto.GetAllApiTokenRes{}

			w.WriteHeader(200)
			response.Message = "ok"
			response.Tokens = ret
			json.NewEncoder(w).Encode(response)
		}
	})))
}
