package dto

type SignInDto struct {
	Username string "json:\"username\" binding:\"required\""
}
