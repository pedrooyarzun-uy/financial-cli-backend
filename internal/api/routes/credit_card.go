package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/middlewares"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewCreditCardRoutes(mux *http.ServeMux, s services.CreditCardService) {
	addCreditCard(mux, s)
}

func addCreditCard(mux *http.ServeMux, s services.CreditCardService) {
	mux.Handle("/credit-card/add", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			userIdVal := r.Context().Value(middlewares.UserID)

			var userId int

			if userIdVal != nil {
				userId = userIdVal.(int)
			}

			var body dto.AddCreditCardReq

			err := json.NewDecoder(r.Body).Decode(&body)

			if err != nil {
				w.WriteHeader(402)
				w.Write([]byte("Bad request"))
				return
			}

			body.OwnerID = userId

			err = s.Add(body)

			response := dto.GenericRes{}

			if err != nil {
				w.WriteHeader(400)
				response.Message = err.Error()
			} else {
				w.WriteHeader(200)
				response.Message = "ok"
			}

			json.NewEncoder(w).Encode(response)
		}
	})))
}
