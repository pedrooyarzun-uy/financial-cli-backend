package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/helpers"
)

func NewUserRoutes(mux *http.ServeMux) {
	signUp(mux)
	signIn(mux)
}

func signUp(mux *http.ServeMux) {
	mux.HandleFunc("/user/sign-up", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var body dto.SignUpReq

			err := json.NewDecoder(r.Body).Decode(&body)

			if err != nil {
				w.WriteHeader(402)
				w.Write([]byte("Something went wrong"))
				return
			}

			if body.Name == "" || body.Email == "" || body.Password == "" {
				w.WriteHeader(400)
				w.Write([]byte("Bad Request"))
				return
			}

			if !helpers.ValidateEmail(body.Email) {
				w.WriteHeader(422)
				w.Write([]byte("Wrong email"))
				return
			}

			//TODO: Sign Up - User (Service)
		}

		w.WriteHeader(405)

		w.Write([]byte("Method not allowed"))
	})
}

func signIn(mux *http.ServeMux) {
	mux.HandleFunc("/user/sign-in", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var body dto.SignInReq

			err := json.NewDecoder(r.Body).Decode(&body)

			if err != nil {
				w.WriteHeader(402)
				w.Write([]byte("Something went wrong"))
				return
			}

			if body.Email == "" || body.Password == "" {
				w.WriteHeader(400)
				w.Write([]byte("Bad Request"))
				return
			}

			if !helpers.ValidateEmail(body.Email) {
				w.WriteHeader(422)
				w.Write([]byte("Wrong email"))
				return
			}

			//TODO: Sign In - User (Service)
		}

		w.WriteHeader(405)

		w.Write([]byte("Method not allowed"))
	})
}
