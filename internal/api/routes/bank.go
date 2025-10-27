package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewBankRoutes(mux *http.ServeMux, s services.BankService) {

}

func getAllBanks(mux *http.ServeMux, s services.BankService) {
	mux.HandleFunc("/bank/get-all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			ret := s.GetAll()

			response := dto.GetAllBankRes{}

			w.WriteHeader(200)
			response.Message = "ok"
			response.Banks = ret
			json.NewEncoder(w).Encode(response)

		}
	})
}
