package models

import "time"

type Item struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;default:current_timestamp();"`
	UpdateAt  time.Time `gorm:"column:updated_at;not null;default:current_timestamp();"`

	Id         uint32   `gorm:"column:id;auto_increment;primary_key"`
	Name       string   `gorm:"column:name;type:varchar(255)"`
	Amount     float32  `gorm:"column:amount;not null"`
	CategoryId uint32   `gorm:"column:category_id;"`
	Category   Category `gorm:"ForeignKey:CategoryId;"`
}
