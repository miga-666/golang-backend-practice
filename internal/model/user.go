package model

import "time"

type User struct {
	Email    string    `gorm:"primaryKey;size:50"  json:"email"`
	Password string    `gorm:"not null;size:255"   json:"-"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
