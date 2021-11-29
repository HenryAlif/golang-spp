package models

type User struct {
	Id       uint
	Name     string
	Classes  string
	Code     string
	Email    string `gorm:"unique"`
	Password []byte
}
