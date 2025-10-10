package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/helpers"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func NewUserRoutes(mux *http.ServeMux, s services.UserService) {
	signUp(mux, s)
	signIn(mux)
}

func signUp(mux *http.ServeMux, s services.UserService) {
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

			//Sign up call
			err = s.SignUp(body)

			//Response
			response := dto.SignUpRes{}

			if err == nil {
				w.WriteHeader(200)
				response.Message = "ok"
				json.NewEncoder(w).Encode(response)
				return
			}

			switch err {
			case services.ErrUserAlreadyExists:
				w.WriteHeader(409)
				response.Message = services.ErrUserAlreadyExists.Error()
			case services.ErrUserCreationFailed:
				w.WriteHeader(400)
				response.Message = services.ErrUserCreationFailed.Error()
			}

			json.NewEncoder(w).Encode(response)
			return

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
