package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string     `json:"id" gorm:"type:string; primary_key;"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func (model *Base) BeforeCreate(tx *gorm.DB) error {
	//id := uuid.NewString()
	//tx.Statement.SetColumn("ID", id)
	return nil
}
