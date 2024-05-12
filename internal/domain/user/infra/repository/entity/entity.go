package entity

import "time"

type User struct {
	Id        string `gorm:"primaryKey;not null;size:36;default:null"`
	Name      string `gorm:"size:255;"`
	Email     string `gorm:"size:255;index:idx_email,unique"`
	Password  string `gorm:"size:72;"`
	CreatedAt time.Time
}
