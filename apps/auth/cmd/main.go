package main

import (
	"backend/apps/auth/internal/model"
	"backend/apps/auth/internal/repository"
	"backend/apps/auth/internal/server"
	"backend/pkg/common/cache"
	"github.com/joho/godotenv"
	"backend/pkg/common/db"
	"go.uber.org/zap"
	"log"
	"os"
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

	redisAddr := os.Getenv("CACHE_ADDRESS") + ":" + os.Getenv("CACHE_PORT")
	redisPassword := os.Getenv("CACHE_PASSWORD")
	redisCache := cache.NewRedisCache(redisAddr, redisPassword, 0)

	authStorage := repository.NewAuthRepo(dbConn, redisCache)

	server.StartServer(&zap.Logger{}, authStorage)
}