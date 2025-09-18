package handlers

import (
	"backend/apps/auth/internal/model"
	"backend/apps/auth/internal/service"
	auth "backend/pkg/gen/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	service service.Auther
}

func NewAuthHandler(service service.Auther) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(ctx context.Context, req *auth.RegisterRequest) (*emptypb.Empty, error) {
	user := model.RegisterUserRequest{
		Login: req.Username,
		Password: req.Password,
	}

	if err := h.service.Register(ctx, user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	creds := model.LoginRequest{
		Login: req.Username,
		Password: req.Password,
	}

	loginResponse, err := h.service.Login(ctx, creds)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "login failed: %v", err)
	}

	return &auth.AuthResponse{
		Token:     loginResponse.AccessToken,
		ExpiresAt: loginResponse.ExpiresAt,
	}, nil
}

func (h *AuthHandler) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	valid, err := h.service.ValidateToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &auth.ValidateResponse{
		Valid: valid.Valid,
		User:  valid.UserID,
	}, nil
}