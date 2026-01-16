package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/middlewares"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewTransactionRoutes(mux *http.ServeMux, s services.TransactionService) {
	add(mux, s)
	getTotalsByCategory(mux, s)
	getTransactionsByDetail(mux, s)
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

			userIdVal := r.Context().Value(middlewares.UserID)

			var userId int

			if userIdVal != nil {
				userId = userIdVal.(int)
			}

			response := dto.GetTotalsByCategoryRes{}

			w.WriteHeader(200)
			response.Message = "ok"
			response.Totals = s.GetTotalsByCategory(userId)
			json.NewEncoder(w).Encode(response)
		}
	})))
}

func getTransactionsByDetail(mux *http.ServeMux, s services.TransactionService) {
	mux.Handle("/transaction/get-all-by-detail", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			//Get user id
			rawUserId := r.Context().Value(middlewares.UserID)

			var userId int

			if rawUserId != nil {
				userId = rawUserId.(int)
			}

			//Get from and to
			rawFrom := r.URL.Query().Get("from")
			rawTo := r.URL.Query().Get("to")

			from, _ := time.Parse("2006-01-02", rawFrom)
			to, _ := time.Parse("2006-01-02", rawTo)

			// Get category and subcategory
			rawCategory := r.URL.Query().Get("category")
			rawSubcategory := r.URL.Query().Get("subcategory")

			category, _ := strconv.Atoi(rawCategory)
			subcategory, _ := strconv.Atoi(rawSubcategory)

			transactions, err := s.GetTransactionsByDetail(userId, from, to, category, subcategory)

			response := dto.GetTransactionsByDetailRes{}

			if err != nil {
				w.WriteHeader(400)
				response.Message = err.Error()
				response.Transactions = []dto.TransactionByDetail{}
			} else {
				w.WriteHeader(200)
				response.Message = "ok"
				response.Transactions = transactions
			}

			json.NewEncoder(w).Encode(response)
		}
	})))
}
