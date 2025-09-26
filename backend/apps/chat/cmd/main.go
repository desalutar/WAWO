package main

import (
	"backend/apps/chat/internal/model"
	"backend/apps/chat/internal/repository"
	"backend/apps/chat/internal/server"
	"backend/pkg/common/cache"
	"backend/pkg/common/db"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load()

	dbConn, err := db.NewPostgresConnection()
	if err != nil {
		log.Fatal("failed to connect DB: ", )
	}

	err = db.RunAutoMigrate(dbConn, model.Dialog{})
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	} 

	redisAddr := os.Getenv("CACHE_ADDRESS") + ":" + os.Getenv("CACHE_PORT")
	redisPassword := os.Getenv("CACHE_PASSWORD")
	redisCache := cache.NewRedisCache(redisAddr, redisPassword, 0)

	chatStorage := repository.NewChatRepo(dbConn, redisCache)

	server.StartChatServer(&zap.Logger{}, chatStorage)
}