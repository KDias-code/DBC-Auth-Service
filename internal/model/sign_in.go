package model

type SignIn struct {
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}
