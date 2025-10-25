package models

import "github.com/Noeeekr/borm"

type CreateUsers struct {
	Name          string   `json:"name" binding:"required,min=2,max=255" borm:"(CONSTRAINTS, NOT NULL)"`
	Email         string   `json:"email" binding:"required,email" borm:"(CONSTRAINTS, UNIQUE, NOT NULL)"`
	Password      string   `json:"password" binding:"required,min=6" borm:"(CONSTRAINTS, NOT NULL)"`
	InstitutionId int      `json:"institution_id" binding:"required" borm:"(NAME, institution_id) (CONSTRAINTS, NOT NULL) (FOREIGN KEY, institutions, id)"`
	Role          UserRole `json:"role" binding:"required" borm:"(TYPE, user_roles) (CONSTRAINTS, NOT NULL)"`

	// Perhaps refactor into a StudentUser table if grows
	Segment *UserSegments `json:"segment" binding:"required" borm:"(TYPE, user_segments)"`
}

type Users struct {
	ID
	DefaultFields
	CreateUsers

	ProfilePicture string `json:"profile_picture" borm:"(NAME, profile_picture) (CONSTRAINTS, DEFAULT 'default_user_pfp')"`
}

func CreateUser(name, email, password string, institutionId int, role UserRole, segment *UserSegments) *CreateUsers {
	return &CreateUsers{
		Name:          name,
		Email:         email,
		Password:      password,
		InstitutionId: institutionId,
		Role:          role,
		Segment:       segment,
	}
}

var TableUsers *borm.TableRegistry = EnvironmentDatabase.RegisterTable(Users{}).
	NeedRoles(TypeUserRole, TypeUserSegment).
	NeedTables(TableInstitutions)
