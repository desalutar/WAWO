package handlers

import (
	"backend/apps/auth/internal/model"
	"backend/apps/auth/internal/service/mocks"
	auth "backend/pkg/gen/proto"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	mocksService := mocks.NewMockAuther(ctrl)

	h := NewAuthHandler(mocksService)

	req := &auth.RegisterRequest{Username: "admin", Password: "admin"}
	mocksService.EXPECT().Register(gomock.Any(), model.RegisterUserRequest{
		Login:    "admin",
		Password: "admin",
	}).Return(nil)

	res, err := h.Register(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}

	require.NoError(t, err)
	require.IsType(t, &emptypb.Empty{}, res)
}