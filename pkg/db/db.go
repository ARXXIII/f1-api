package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var Conn *pgx.Conn

func InitDB(ctx context.Context) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	user := os.Getenv("POSGRES_USER")
	password := os.Getenv("POSGRES_PASSWORD")
	host := os.Getenv("POSGRES_HOST")
	port := os.Getenv("POSGRES_PORT")
	dbname := os.Getenv("POSGRES_DBNAME")

	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		log.Fatal("One or more required database connection parameters are missing")
	}

	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	var err error
	Conn, err = pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Connected to Supabase Postgres")
}
