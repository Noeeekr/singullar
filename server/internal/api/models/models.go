package models

type Signin struct { // FOR JSON
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=2"`
}
