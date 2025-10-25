package models

type Subjects struct {
	ID
	InstitutionId int    `borm:"(NAME, institution_id) (FOREIGN KEY, institutions, id)"`
	SubjectName   string `borm:"(NAME, subject_name)"`
}

var TableSubjects = EnvironmentDatabase.
	RegisterTable(Subjects{}).
	NeedTables(TableInstitutions)
