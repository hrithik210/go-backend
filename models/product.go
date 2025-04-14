package models

import (
	"time"

	"github/lib/pq"

	"github.com/lib/pq"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Description string         `json:"description" gorm:"not null"`
	Price       float64        `json:"price" gorm:"not null"`
	Category    string         `json:"category"`
	IsMainCard  bool           `json:"isMainCard" gorm:"default:false"`
	Colors      pq.StringArray `json:"colors" gorm:"type:text[]"`
	Sizes       pq.StringArray `json:"sizes" gorm:"type:text[]"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
}
