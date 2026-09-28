package models

import "time"

type TemplateColor struct {
	Id          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `gorm:"type:varchar(255);not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	UserID      *uint16    `gorm:"default:null;constraint:OnDelete:SET NULL;" json:"user_id"`
	User        User       `gorm:"foreignKey:UserID;references:Id" json:"user,omitempty"`
	Colors      []Color    `gorm:"many2many:template_color_relations;constraint:OnDelete:CASCADE;" json:"colors,omitempty"`
	Categories  []Category `gorm:"many2many:template_category_relations;constraint:OnDelete:CASCADE;" json:"categories,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
