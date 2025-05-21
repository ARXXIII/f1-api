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
	constructorRepo := repository.NewConstructorRepository()

	driverService := service.NewDriverService(driverRepo)
	constructorService := service.NewConstructorService(constructorRepo)

	driverHandler := handler.NewDriverHandler(ctx, driverService)
	constructorHandler := handler.NewConstructorHandler(ctx, constructorService)

	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/driver", driverHandler.GetDriver)
	http.HandleFunc("/driver/", driverHandler.GetDriverByID)
	http.HandleFunc("/constructor", constructorHandler.GetConstructor)
	http.HandleFunc("/constructor/", constructorHandler.GetConstructorByID)

	port := os.Getenv("PORT")
	log.Printf("Server is running on http://localhost:%s/health", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
