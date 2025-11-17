package models

// Request with client populated fields
type CreateSubjectsRequest struct {
	SubjectName string `borm:"(NAME, subject_name)" json:"subject_name" binding:"required"`
}

// Request with server populated fields
type CreateSubjects struct {
	CreateSubjectsRequest
	InstitutionId int `borm:"(NAME, institution_id) (FOREIGN KEY, institutions, id)"`
}

// Request with database populated fields
type Subjects struct {
	ID
	CreateSubjects
}

var TableSubjects = EnvironmentDatabase.
	RegisterTable(Subjects{}).
	NeedTables(TableInstitutions)
