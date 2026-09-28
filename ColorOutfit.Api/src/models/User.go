package models

import "time"

type User struct {
	Id             uint16          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string          `gorm:"type:varchar(255);not null" json:"name"`
	Username       string          `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email          string          `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password       string          `gorm:"type:varchar(255);not null" json:"-"`
	TemplateColors []TemplateColor `gorm:"foreignKey:UserID" json:"template_colors,omitempty"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
