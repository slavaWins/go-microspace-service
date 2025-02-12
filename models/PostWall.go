package models

import "time"

type PostWall struct {
	ID         uint `gorm:"primarykey" json:"id"`
	Created_at time.Time
	Updated_at time.Time
	Content    string
	UserId     uint
	ViewsCount uint
}
