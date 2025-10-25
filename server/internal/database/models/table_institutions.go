package models

import "github.com/Noeeekr/borm"

type CreateInstitutions struct {
	Name     string `binding:"required" borm:"(CONSTRAINTS, NOT NULL)" json:"name"`
	Email    string `binding:"required" json:"email" borm:"(IGNORE)"`
	Password string `binding:"required" json:"password" borm:"(IGNORE)"`
}

type Institutions struct {
	ID
	DefaultFields
	CreateInstitutions
}

var TableInstitutions *borm.TableRegistry = EnvironmentDatabase.RegisterTable(Institutions{})
