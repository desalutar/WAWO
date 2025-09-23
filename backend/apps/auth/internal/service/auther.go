package service

import (
	"backend/apps/auth/internal/model"
	"context"
)

type Auther interface {
	Register(ctx context.Context, user model.RegisterUserRequest) error
	Login(ctx context.Context, userDTO model.LoginRequest) (*model.LoginResponse, error)
	ValidateToken(ctx context.Context, token string) (*model.ValidateTokenResponse, error)
	Logout(id int) error
}