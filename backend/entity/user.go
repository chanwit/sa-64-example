package entity

import (
  "gorm.io/gorm"
  "time"
)

type User struct {
  gorm.Model
  FirstName    string
  LastName     string
  Email        string
  Age          uint8
  BirthDay     time.Time
}