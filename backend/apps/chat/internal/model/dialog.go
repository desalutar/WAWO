package model

import "time"

type Dialog struct {
    ID              uint                    `gorm:"primaryKey"`
    Participants    []DialogParticipant     `gorm:"foreignKey:DialogID"`
    LastMessage     string
    LastUpdated     time.Time
}

type DialogParticipant struct {
    ID          uint `gorm:"primaryKey"`
    DialogID    uint
    UserID      uint
}