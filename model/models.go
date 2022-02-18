package models

type User struct {
	Id        string `json:"id" gorm:"primary_key `
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
}
