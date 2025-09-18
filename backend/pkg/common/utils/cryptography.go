package utils

import (
	"backend/pkg/common/errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AccessToken = iota
	RefreshToken
)

type Token struct {
	AccessTTL     time.Duration `yaml:"access_ttl"`
	RefreshTTL    time.Duration `yaml:"refresh_ttl"`
	AccessSecret  string        `yaml:"access_secret"`
	RefreshSecret string        `yaml:"refresh_secret"`
}

type TokenManager interface {
	CreateToken(userID, role string, ttl time.Duration, kind int) (string, error)
	ParseToken(inputToken string, kind int) (UserClaims, error)
}

type TokenJWT struct {
	AccessSecret 	[]byte
	RefreshSecret   []byte
}

func NewTokenJWT(token Token) TokenManager {
	return &TokenJWT{AccessSecret: []byte(token.AccessSecret), RefreshSecret: []byte(token.RefreshSecret)}
}

type UserClaims struct {
	ID     string `json:"uid"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type UserFromClaims struct {
	ID     int
	Role   int
}

func (o *TokenJWT) CreateToken(userID, role string, ttl time.Duration, kind int) (string, error) {
	claims := UserClaims {
		ID: 				userID,
		Role: 				role,
		RegisteredClaims: 	jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var secret []byte
	switch kind {
	case AccessToken:
		secret = o.AccessSecret
	case RefreshToken:
		secret = o.RefreshSecret
	default:
		return "", errors.TokenTypeError
	}

	return token.SignedString(secret)
}

func (o *TokenJWT) ParseToken(inputToken string, kind int) (UserClaims, error) {
	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		var secret []byte
		switch kind {
		case AccessToken:
			secret = o.AccessSecret
		case RefreshToken:
			secret = o.RefreshSecret
		default:
			return "", errors.TokenTypeError
		}
		_ = secret

		return secret, nil
	})

	if err != nil {
		return UserClaims{}, err
	}

	if !token.Valid {
		return UserClaims{}, fmt.Errorf("not valid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return UserClaims{}, fmt.Errorf("error get user claims from token")
	}

	return UserClaims{
		ID:               claims["uid"].(string),
		Role:             claims["role"].(string),
		RegisteredClaims: jwt.RegisteredClaims{},
	}, nil
}