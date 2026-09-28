package models

import "time"

type Category struct {
	Id        uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string          `gorm:"type:varchar(255);not null" json:"name"`
	Templates []TemplateColor `gorm:"many2many:template_category_relations;constraint:OnDelete:CASCADE;" json:"templates,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
