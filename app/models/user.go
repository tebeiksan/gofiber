package models

type User struct {
	Id    string `gorm:"primaryKey" json:"id"`
	Email string `gorm:"unique,index" json:"email"`
	Name  string `json:"name"`
}
