package model

import "time"

type Dialog struct {
    ID           uint      `gorm:"primaryKey"`
    ParticipantIDs []uint   `gorm:"type:integer[]"`
    LastMessage  string
    LastUpdated  time.Time
}