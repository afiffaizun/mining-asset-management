package entity

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	AssetCode string    `gorm:"type:varchar(30);uniqueIndex" json:"asset_code"`
	AssetName string    `gorm:"type:varchar(100)" json:"asset_name"`
	Status    string    `gorm:"type:varchar(20)" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
