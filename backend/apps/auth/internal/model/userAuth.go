package model

type UserAuth struct {
    ID       uint    `gorm:"primaryKey"`
    Login    string  `gorm:"size:50;unique;not null"`
    Password string  `gorm:"size:100;not null"`
    Role     *string `json:"role" gorm:"column:role;type:varchar(255)"`
}

type LoginRequest struct {
    Login string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiresAt    int64
}


type RegisterUserRequest struct {
	Login     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type ValidateTokenResponse struct {
	Valid bool
	UserID string
	ErrorMessage string
}