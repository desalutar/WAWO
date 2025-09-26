package server

import (
	"backend/apps/chat/internal/handlers"
	"backend/apps/chat/internal/repository"
	"backend/apps/chat/internal/service"
	chatpb "backend/pkg/gen/chat/proto"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func StartChatServer(logger *zap.Logger, storage repository.Chater) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Printf("Failed to listen chatServer: %v", err)
	}

	grpcServer := grpc.NewServer()

	chatService := service.NewChatService(storage)

	chatpb.RegisterChatServiceServer(grpcServer, handlers.NewChatHandlers(chatService))	

	reflection.Register(grpcServer)

	log.Println("Chat_gRPC_Server запущен на :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("ошибка запуска Chat_gRPC_Server: %v", err)
	}
}