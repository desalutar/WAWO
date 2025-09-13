package run

import (
	"context"
	"fmt"
	"log"

	"backend/pkg/db"
)

func Run(ctx context.Context) {

	db.Init()
	if err := db.DB.AutoMigrate(&db.User{}, &db.Session{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("Auth service is running...")

	<-ctx.Done()

	fmt.Println("Shutting down gracefully...")
	sqlDB, err := db.DB.DB()
	if err == nil {
		sqlDB.Close()
	}
	fmt.Println("Shutdown complete")
}
