package service

import (
	"backend/apps/chat/internal/model"
	"backend/apps/chat/internal/repository"
	"context"
)

type ChatServicer interface {
	GetDialogs(ctx context.Context, userID uint) ([]model.Dialog, error)
}

type ChatService struct {
	repo repository.Chater
}

func NewChatService(r repository.Chater) *ChatService {
	return &ChatService{
		repo: r,
	}
}

func (c *ChatService) GetDialogs(
	ctx context.Context, 
	userID uint) ([]model.Dialog, error) {
		return c.repo.GetDialogs(userID)
}