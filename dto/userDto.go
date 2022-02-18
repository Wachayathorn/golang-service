package dto

type CreateUserDto struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Username  string `json:"username" binding:"required"`
}
