package main

import (
	"log"
	"net/http"

	"github.com/ARXXIII/f1-api/internal/handler"
	"github.com/ARXXIII/f1-api/internal/repository"
	"github.com/ARXXIII/f1-api/internal/service"
)

func main() {
	// Конфигурация
	supabaseURL := "https://your-project-id.supabase.co/rest/v1"
	supabaseAPIKey := "your-supabase-api-key"

	// DI — dependency injection
	driverRepo := repository.NewDriverRepository(supabaseURL, supabaseAPIKey)
	driverService := service.NewDriverService(driverRepo)
	driverHandler := handler.NewDriverHandler(driverService)

	// Маршруты
	http.HandleFunc("/driver", driverHandler.GetDrivers)

	log.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
