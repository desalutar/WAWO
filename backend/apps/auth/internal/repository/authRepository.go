package repository

import (
	"backend/apps/auth/internal/model"
	"backend/pkg/common/cache"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Auther interface {
	GetByEmail(ctx context.Context, email string) (model.UserAuth, error)
	GetByID(ctx context.Context, id string) (model.UserAuth, error)
	CreateUser(ctx context.Context, user model.UserAuth) error
	SaveAccessToken(ctx context.Context, userID string, accessToken string, expires time.Duration) error
	Logout(id int) error
}

type Auth struct {
	db      *gorm.DB
	cache  	cache.Cacher
}

func NewAuthRepo(db *gorm.DB, cache cache.Cacher) *Auth {
	return &Auth{
		db: 	 db,
		cache: 	 cache,
	}
}

func (a *Auth) CreateUser(ctx context.Context, user model.UserAuth) error {
	return a.db.WithContext(ctx).Create(&user).Error
}

func (a *Auth) GetByEmail(ctx context.Context, email string) (model.UserAuth, error) {
	var user model.UserAuth
	err := a.db.WithContext(ctx).Where("login = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.UserAuth{}, fmt.Errorf("user not found")
	}
	return user, err
}

func (a *Auth) SaveAccessToken(ctx context.Context, userID string, accessToken string, expires time.Duration) error {
	key := "access_token:" + userID
	return a.cache.Set(ctx, key, accessToken, expires)
}

// TODO: Доделать
func (a *Auth) Logout(userID int) error {
	// Удаление access_token из кэша по ключу
	// key := "access_token:" + userID
	// return a.cache.(context.Background(), key)
	return nil
}

func (r *Auth) GetByID(ctx context.Context, id string) (model.UserAuth, error) {
	var user model.UserAuth
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return model.UserAuth{}, err
	}
	return user, nil
}