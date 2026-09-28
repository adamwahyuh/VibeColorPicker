package models

import "time"

type Color struct {
	Id        uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Hex       string          `gorm:"type:varchar(10);not null" json:"hex"`
	Templates []TemplateColor `gorm:"many2many:template_color_relations;constraint:OnDelete:CASCADE;" json:"templates,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
