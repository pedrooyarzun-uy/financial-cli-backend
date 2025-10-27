package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewAccountRoutes(mux *http.ServeMux, s services.AccountService) {
	create(mux, s)
}

func create(mux *http.ServeMux, s services.AccountService) {
	mux.HandleFunc("/account/create", func(w http.ResponseWriter, r *http.Request) {
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
	})
}
