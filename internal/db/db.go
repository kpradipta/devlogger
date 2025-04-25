package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() error {
	dsn := os.Getenv("DATABASE_URL") // Railway provides this env var
	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	log.Println("Connected to db")
	return nil
}
