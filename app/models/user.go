package models

type User struct {
	Id    string `gorm:"primaryKey"`
	Email string `gorm:"index"`
	Name  string
}
