package main

import (
	"log"
	"backend/pkg/common/db"
	"github.com/joho/godotenv"
	"backend/apps/auth/internal/model"
)

func main() {
	_ = godotenv.Load()

	dbConn, err := db.NewPostgresConnection()
	if err != nil {
		log.Fatal("failed to connect DB: ", )
	}

	err = db.RunAutoMigrate(dbConn, model.UserAuth{})
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	} 

}