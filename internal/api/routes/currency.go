package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewCurrencyRoutes(mux *http.ServeMux, s services.CurrencyService) {
	getAll(mux, s)
}

func getAll(mux *http.ServeMux, s services.CurrencyService) {
	mux.HandleFunc("/currency/get-all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			ret := s.GetAll()

			response := dto.GetAllCurrencyRes{}

			w.WriteHeader(200)
			response.Message = "ok"
			response.Currencies = ret
			json.NewEncoder(w).Encode(response)

		}
	})
}
