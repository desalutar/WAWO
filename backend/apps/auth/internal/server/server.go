package server

import (
	"backend/apps/auth/internal/handlers"
	"backend/apps/auth/internal/repository"
	"backend/apps/auth/internal/service"
	"backend/pkg/common/utils"
	authpb "backend/pkg/gen/proto"
	"log"
	"net"
	"os"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer(logger *zap.Logger, storage repository.Auther) {
	tokenConfig := utils.Token{
		AccessTTL:     time.Minute * 15,
		RefreshTTL:    time.Hour * 24 * 7,
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Printf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authService := service.NewAuthService(storage, utils.NewTokenJWT(tokenConfig))
	
	authpb.RegisterAuthServiceServer(grpcServer, handlers.NewAuthHandler(authService))

	reflection.Register(grpcServer)

	log.Println("gRPC-Сервер запущен на :50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("ошибка запуска gRPC-сервера: %v", err)
	}
}