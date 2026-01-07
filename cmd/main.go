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
	tr := repositories.NewTransactionRepository(db.DB)
	cr := repositories.NewCurrencyRepository(db.DB)
	br := repositories.NewBankRepository(db.DB)
	catr := repositories.NewCategoryRepository(db.DB)
	atr := repositories.NewApiTokenRepository(db.DB)
	sr := repositories.NewSubcategoryRepository(db.DB)

	//Services
	us := services.NewUserService(ur)
	as := services.NewAccountService(ar)
	ts := services.NewTransactionRepository(tr, ar)
	cs := services.NewCurrencyService(cr)
	bs := services.NewBankService(br)
	cats := services.NewCategoryService(catr)
	ats := services.NewApiTokenService(atr)
	ss := services.NewSubcategoryService(sr)

	routes.NewUserRoutes(mux, us)
	routes.NewAccountRoutes(mux, as)
	routes.NewTransactionRoutes(mux, ts)
	routes.NewCurrencyRoutes(mux, cs)
	routes.NewBankRoutes(mux, bs)
	routes.NewCategoryRoutes(mux, cats)
	routes.NewApiToken(mux, ats)
	routes.NewSubcategoryRoutes(mux, ss)

	fmt.Println("Server started at :8080")

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
