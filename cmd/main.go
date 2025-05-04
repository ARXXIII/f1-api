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
	driverService := service.NewDriverService(driverRepo)
	driverHandler := handler.NewDriverHandler(ctx, driverService)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/driver", driverHandler.GetDrivers)
	http.HandleFunc("/driver/", driverHandler.GetDriverByID)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
