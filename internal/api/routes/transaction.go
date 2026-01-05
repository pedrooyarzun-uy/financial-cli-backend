package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/middlewares"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewTransactionRoutes(mux *http.ServeMux, s services.TransactionService) {
	add(mux, s)
	getTotalsByCategory(mux, s)
}

func add(mux *http.ServeMux, s services.TransactionService) {
	mux.Handle("/transaction/add", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body dto.AddTransactionReq

			err := json.NewDecoder(r.Body).Decode(&body)

			if err != nil {
				w.WriteHeader(402)
				w.Write([]byte("Something went wrong"))
				return
			}

			if body.Amount == 0 || body.Account == 0 || body.Currency == 0 || body.Type == 0 {
				w.WriteHeader(400)
				w.Write([]byte("Bad Request"))
				return
			}

			err = s.Add(body)

			response := dto.GenericRes{}

			if err == nil {
				w.WriteHeader(200)
				response.Message = "ok"
				json.NewEncoder(w).Encode(response)
				return
			}

			switch err {
			case services.ErrTransactionNotCorrectCurrency:
				w.WriteHeader(409)
				response.Message = err.Error()
			}

			json.NewEncoder(w).Encode(response)

		}
	})))
}

func getTotalsByCategory(mux *http.ServeMux, s services.TransactionService) {
	mux.Handle("/transaction/get-totals-by-category", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			response := dto.GetTotalsByCategoryRes{}

			w.WriteHeader(200)
			response.Message = "ok"
			response.Totals = s.GetTotalsByCategory()
			json.NewEncoder(w).Encode(response)
		}
	})))
}
