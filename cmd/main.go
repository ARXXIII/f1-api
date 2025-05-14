package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ARXXIII/f1-api/internal/handler"
	"github.com/ARXXIII/f1-api/internal/repository"
	"github.com/ARXXIII/f1-api/internal/service"
	"github.com/ARXXIII/f1-api/pkg/db"
)

func main() {
	ctx := context.Background()

	db.InitDB(ctx)
	defer db.Conn.Close(ctx)

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

	port := os.Getenv("PORT")
	log.Printf("Server is running on http://localhost:%s/health", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
