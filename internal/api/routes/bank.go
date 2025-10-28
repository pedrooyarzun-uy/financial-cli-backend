package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/middlewares"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewBankRoutes(mux *http.ServeMux, s services.BankService) {
	getAllBanks(mux, s)
}

func getAllBanks(mux *http.ServeMux, s services.BankService) {
	mux.Handle("/bank/get-all", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			ret := s.GetAll()

			response := dto.GetAllBankRes{}

			w.WriteHeader(200)
			response.Message = "ok"
			response.Banks = ret
			json.NewEncoder(w).Encode(response)
		}
	})))
}
