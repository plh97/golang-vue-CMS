package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 用于替代 gorm.Model，统一 JSON 格式为 camelCase
type BaseModel struct {
	// ID 统一为小写的 id
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
