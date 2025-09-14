package model

type UserAuth struct {
    ID       uint   `gorm:"primaryKey"`
    Login    string `gorm:"size:50;unique;not null"`
    Password string `gorm:"size:100;not null"`
}
