package main

import (
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/routes"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/server"
)

func main() {
	srv, mux := server.NewServer(":8080")
	//Load routes
	routes.NewUserRoutes(mux)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
