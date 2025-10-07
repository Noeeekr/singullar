package models

import "github.com/Noeeekr/borm"

type CreateClassRequest struct {
	Name    string       `json:"className" binding:"required"`
	Segment UserSegments `json:"segment" binding:"required"`
	Series  string       `json:"series" binding:"required"`

	StudentsIds []int `json:"students" binding:"required"`
	TeacherId   int   `json:"teacher_id" binding:"required" borm:"(NAME, teacher_id) (CONSTRAINTS, NOT NULL) (FOREIGN KEY, institutions, id)"`
}
type CreateClasses struct {
	Name    string       `json:"name" binding:"required"`
	Segment UserSegments `json:"segment" binding:"required" borm:"(TYPE, user_segments) (CONSTRAINTS, NOT NULL)"`
	Series  string       `json:"series" binding:"required"`

	InstitutionId int `json:"institution_id" binding:"required" borm:"(NAME, institution_id) (CONSTRAINTS, NOT NULL) (FOREIGN KEY, institutions, id)"`
	TeacherId     int `json:"teacher_id" binding:"required" borm:"(NAME, teacher_id) (CONSTRAINTS, NOT NULL) (FOREIGN KEY, users, id)"`
}
type Classes struct {
	ID
	DefaultFields
	CreateClasses
}

var TableClasses *borm.TableRegistry = EnvironmentDatabase.RegisterTable(Classes{}).
	NeedTables(TableUsers)
