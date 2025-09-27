package repository

import (
	"backend/apps/chat/internal/model"
	"backend/pkg/common/cache"
	"log"

	"gorm.io/gorm"
)

type Chater interface {
    GetDialogs(userID uint) ([]model.Dialog, error)
    GetMessages(dialogID uint) ([]model.Message, error)    
    SendMessage(msg model.Message) error
    CreateDialog(participantIDs []uint) (model.Dialog, error)
}

type Chat struct {
	db 		*gorm.DB
	cache 	cache.Cacher
}

func NewChatRepo(db *gorm.DB, cache cache.Cacher) *Chat {
	return &Chat{
		db: db,
		cache: cache,
	}
}

// func (c *Chat) GetDialogs(userID uint) ([]model.Dialog, error) {
// 	var dialogs []model.Dialog
// 	err := c.db.
// 		Where("? = ANY(participant_ids)", userID).
// 		Order("last_updated desc").
// 		Find(&dialogs).Error
// 	if err != nil {
// 		log.Println("Error in GetDialogs")
// 		return nil, err 
// 	}
	
// 	return dialogs , nil
// }

func (c *Chat) GetDialogs(userID uint) ([]model.Dialog, error) {
    var dialogs []model.Dialog
	err := c.db.
		Preload("Participants", "user_id != ?", userID).
		Joins("JOIN dialog_participants dp ON dp.dialog_id = dialogs.id").
		Where("dp.user_id = ?", userID).
		Order("last_updated DESC").
		Find(&dialogs).Error

    if err != nil {
        log.Println("Error in GetDialogs:", err)
        return nil, err
    }

    return dialogs, nil
}

func (c *Chat) GetMessages(dialogID uint) ([]model.Message, error) {
	return []model.Message{}, nil
}

func (c *Chat) SendMessage(msg model.Message) error {
	return nil
}

func (c *Chat) CreateDialog(participantIDs []uint) (model.Dialog, error) {
	return model.Dialog{}, nil
}