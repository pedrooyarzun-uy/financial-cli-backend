package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/routes"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/db"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/server"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/services"
)

func main() {
	srv, mux := server.NewServer(":8080")
	//Load routes

	if os.Getenv("APPLICATION_ENV") == "" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal("Couldn't load .env file", err)
		}
	}
	db.Init()

	//Repositories
	ur := repositories.NewUserRepository(db.DB)
	ar := repositories.NewAccountRepository(db.DB)

	//Services
	us := services.NewUserService(ur)
	as := services.NewAccountService(ar)

	routes.NewUserRoutes(mux, us)
	routes.NewAccountRoutes(mux, as)

	fmt.Println("Server is starting!")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
