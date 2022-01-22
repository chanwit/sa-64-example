package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `valid:"required~Name cannot be blank"`
	Email     string `gorm:"uniqueIndex" valid:"email"`
	StudentID string `valid:"matches(^[BMD]\\d{7}$)"`
	Password  string

	// 1 user เป็นเจ้าของได้หลาย video
	Videos []Video `gorm:"foreignKey:OwnerID"`
	// 1 user เป็นเจ้าของได้หลาย playlist
	Playlists []Playlist `gorm:"foreignKey:OwnerID"`
}

type Video struct {
	gorm.Model
	Name string
	Url  string `gorm:"uniqueIndex"`
	// OwnerID ทำหน้าที่เป็น FK
	OwnerID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Owner       User         `gorm:"references:id"`
	WatchVideos []WatchVideo `gorm:"foreignKey:VideoID"`
}

type Playlist struct {
	gorm.Model
	Title string
	// OwnerID ทำหน้าที่เป็น FK
	OwnerID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Owner       User         `gorm:"references:id"`
	WatchVideos []WatchVideo `gorm:"foreignKey:PlaylistID"`
}

type Resolution struct {
	gorm.Model
	Value       string
	WatchVideos []WatchVideo `gorm:"foreignKey:ResolutionID"`
}

type WatchVideo struct {
	gorm.Model
	WatchedTime time.Time

	// ResolutionID ทำหน้าที่เป็น FK
	ResolutionID *uint
	Resolution   Resolution `gorm:"references:id"`

	// PlaylistID ทำหน้าที่เป็น FK
	PlaylistID *uint
	Playlist   Playlist `gorm:"references:id"`

	// VideoID ทำหน้าที่เป็น FK
	VideoID *uint
	Video   Video `gorm:"references:id"`
}
