package handlers

import (
	"backend/apps/chat/internal/service"
	chat "backend/pkg/gen/chat/proto"
	"context"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type ChatHandler struct {
	chat.UnimplementedChatServiceServer
	service service.ChatServicer
}

func NewChatHandlers(service *service.ChatService) *ChatHandler {
	return &ChatHandler{
		service: service,
	}
}

func (h *ChatHandler) GetDialogs(ctx context.Context, req *chat.GetDialogsRequest) (*chat.GetDialogsResponse, error) {
	dialogs, err := h.service.GetDialogs(ctx, uint(req.UserId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get dialogs: %v", err)
	}

	var protoDialogs []*chat.DialogSummary
	for _, d := range dialogs {
		protoDialogs = append(protoDialogs, &chat.DialogSummary{
			DialogId: 		uint32(d.ID),
			ParticipantIds: uintSliceToUint32(d.ParticipantIDs),
			LastMessage: 	d.LastMessage,
			LastUpdated: 	d.LastUpdated.Unix(),
		})
	}

	return &chat.GetDialogsResponse{Dialogs: protoDialogs}, nil
}

func uintSliceToUint32(s []uint) []uint32 {
    res := make([]uint32, len(s))
    for i, v := range s {
        res[i] = uint32(v)
    }
    return res
}