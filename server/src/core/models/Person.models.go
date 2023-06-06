package models

import "time"

type Person struct {
	ID           uint      `gorm:"column:Id" json:"id"`
	Name         string    `gorm:"column:Name;not null" json:"name"`
	LastName     string    `gorm:"column:LastName;not null" json:"lastName"`
	CreationDate time.Time `gorm:"column:CreationDate;not null;default:CURRENT_TIMESTAMP" json:"creationDate"`
	Active       bool      `gorm:"column:Active;default:true;not null" json:"active"`
}
