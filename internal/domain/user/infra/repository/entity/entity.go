package entity

import "time"

type User struct {
	Id        string `gorm:"primaryKey;size:36;"`
	Name      string `gorm:"size:255;"`
	Email     string `gorm:"size:255;index:idx_email,unique"`
	Password  string `gorm:"size:72;"`
	CreatedAt time.Time
}
