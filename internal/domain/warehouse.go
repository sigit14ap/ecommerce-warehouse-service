package domain

import "time"

type Warehouse struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"not null;size:255" json:"name"`
	Location  string    `gorm:"not null;size:255" json:"location"`
	IsActive  bool      `gorm:"not null;" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
