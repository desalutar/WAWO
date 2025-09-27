package model

import "time"

type Message struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    DialogID  uint      `gorm:"index" json:"dialog_id"`
    SenderID  uint      `json:"sender_id"`
    Text      string    `json:"text"`
    Timestamp time.Time `json:"timestamp"`
}