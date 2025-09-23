package service

import (
	"backend/apps/auth/internal/model"
	"backend/apps/auth/internal/repository"
	"backend/pkg/common/utils"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type Auth struct {
	repo         repository.Auther
	tokenManager utils.TokenManager
}

func NewAuthService(repo repository.Auther, tokenManager utils.TokenManager) *Auth {
	return &Auth{
		repo:         repo,
		tokenManager: tokenManager,
	}
}

func (a *Auth) Register(ctx context.Context, userDTO model.RegisterUserRequest) error {
	existingUser, err := a.repo.GetByEmail(ctx, userDTO.Login)
	if err == nil && existingUser.ID != 0 {
		return fmt.Errorf("user with email %s already exists", userDTO.Login)
	}

	hashedPassword, err := utils.HashPassword(userDTO.Password)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return fmt.Errorf("internal error")
	}

	newUser := model.UserAuth{
		Login:     userDTO.Login,
		Password:  string(hashedPassword),
	}

	if err := a.repo.CreateUser(ctx, newUser); err != nil {
		log.Printf("failed to create user: %v", err)
		return fmt.Errorf("could not create user")
	}
	log.Printf("User registered: email=%s", userDTO.Login)
	return nil
}

func (a *Auth) Login(ctx context.Context, userDTO model.LoginRequest) (*model.LoginResponse, error) {
	user, err := a.repo.GetByEmail(ctx, userDTO.Login)
	if err != nil {
		return &model.LoginResponse{}, fmt.Errorf("invalid email or password")
	}

	if !utils.CheckPassword(user.Password, userDTO.Password) {
		return &model.LoginResponse{}, fmt.Errorf("invalid email or password")
	}

	accessToken, accessExpiresAt, refreshToken, err := a.generateTokens(fmt.Sprintf("%d", user.ID), fmt.Sprintf("%v", user.Role))
	if err != nil {
		return &model.LoginResponse{}, err
	}

	accessTokenExpires := time.Minute * 2
	err = a.repo.SaveAccessToken(ctx, fmt.Sprintf("%d", user.ID), accessToken, accessTokenExpires)
	if err != nil {
		log.Printf("failed to save access token: %v", err)
		return &model.LoginResponse{}, fmt.Errorf("failed to save access token")
	}

	log.Printf("User logged in: email=%s", userDTO.Login)
	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessExpiresAt,
	}, nil
}

func (a *Auth) Logout(id int) error {
	log.Printf("User logged out: id=%d", id)
	return nil
}

func (a *Auth) generateTokens(userID, role string) (accessToken string, accessExpiresAt int64, refreshToken string, err error) {
    accessTTL := time.Minute * 15
    refreshTTL := time.Hour * 24 * 7

    accessToken, err = a.tokenManager.CreateToken(userID, role, accessTTL, utils.AccessToken)
    if err != nil {
        return "", 0, "", fmt.Errorf("internal error")
    }
    refreshToken, err = a.tokenManager.CreateToken(userID, role, refreshTTL, utils.RefreshToken)
    if err != nil {
        return "", 0, "", fmt.Errorf("internal error")
    }

    accessExpiresAt = time.Now().Add(accessTTL).Unix()
    return accessToken, accessExpiresAt, refreshToken, nil
}

func (a *Auth) ValidateToken(ctx context.Context, token string) (*model.ValidateTokenResponse, error) {
	parsedToken, err := a.tokenManager.ParseToken(token, utils.AccessToken)
	if err != nil {
    	return &model.ValidateTokenResponse{}, errors.New("invalid token")
	}

	userID := parsedToken.ID

	user, err := a.repo.GetByID(ctx, userID)
	if err != nil || user.ID == 0 {
		return &model.ValidateTokenResponse{
			Valid:        false,
			ErrorMessage: "user not found",
		}, nil
	}

	return &model.ValidateTokenResponse{
		Valid:  true,
		UserID: userID,
	}, nil
}
