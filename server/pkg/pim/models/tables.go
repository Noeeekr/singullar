package models

type TableName string

const (
	UsersTableName TableName = "user"
)

type Users struct {
	Name     string `json:"name" binding:"required,min=2,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=2"`
	Role     Roles  `json:"role" binding:"required"`
}
