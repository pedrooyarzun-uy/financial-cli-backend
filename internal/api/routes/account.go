package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/middlewares"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewAccountRoutes(mux *http.ServeMux, s services.AccountService) {
	create(mux, s)
}

func create(mux *http.ServeMux, s services.AccountService) {
	mux.Handle("/account/create", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body dto.CreateReq

			err := json.NewDecoder(r.Body).Decode(&body)

			if err != nil {
				w.WriteHeader(402)
				w.Write([]byte("Bad request"))
				return
			}

			err = s.Create(body)

			response := dto.GenericRes{}

			if err == nil {
				w.WriteHeader(200)
				response.Message = "ok"
				json.NewEncoder(w).Encode(response)
				return
			}

			switch err {
			case services.ErrAccountAlreadyExists:
				w.WriteHeader(409)
				response.Message = err.Error()

			}

			json.NewEncoder(w).Encode(response)
		}
	})))

	mux.Handle("/account/get-all", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			userIdVal := r.Context().Value(middlewares.UserID)

			var userId int

			if userIdVal != nil {
				userId = userIdVal.(int)
			}

			res := s.GetAll(userId)

			response := dto.GetAllAccountRes{
				Message:  "ok",
				Accounts: res,
			}

			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response)

		}
	})))
}
