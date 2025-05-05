package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ARXXIII/f1-api/internal/handler"
	"github.com/ARXXIII/f1-api/internal/repository"
	"github.com/ARXXIII/f1-api/internal/service"
	"github.com/ARXXIII/f1-api/pkg/db"
)

func main() {
	ctx := context.Background()

	db.InitDB(ctx)
	defer db.Conn.Close(ctx)

	// var version string
	// if err := db.Conn.QueryRow(ctx, "SELECT version()").Scan(&version); err != nil {
	// 	log.Fatalf("Query failed: %v", err)
	// }
	// log.Println("Postgres version:", version)

	driverRepo := repository.NewDriverRepository()
	teamRepo := repository.NewTeamRepository()

	driverService := service.NewDriverService(driverRepo)
	teamService := service.NewTeamService(teamRepo)

	driverHandler := handler.NewDriverHandler(ctx, driverService)
	teamHandler := handler.NewTeamHandler(ctx, teamService)

	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/driver", driverHandler.GetDrivers)
	http.HandleFunc("/driver/", driverHandler.GetDriverByID)
	http.HandleFunc("/team", teamHandler.GetTeams)
	http.HandleFunc("/team/", teamHandler.GetTeamByID)

	log.Println("Server is running on http://localhost:8080/health")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
