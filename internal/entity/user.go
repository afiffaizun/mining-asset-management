package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	RoleID    uuid.UUID `gorm:"type:uuid" json:"role_id"`
	FullName  string    `gorm:"type:varchar(100)" json:"full_name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password  string    `gorm:"type:text" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	Role      Role      `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}
