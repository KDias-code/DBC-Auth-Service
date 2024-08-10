package model

type SignUp struct {
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
	Iin      string `json:"iin" db:"iin"`
}
