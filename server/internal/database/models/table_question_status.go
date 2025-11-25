package models

// Not implemented yet
type QuestionStatus struct {
	ID
	StatusName int `borm:"(NAME, status_name)"`
}

var TableQuestionStatus = EnvironmentDatabase.
	RegisterTable(QuestionStatus{}).
	NeedTables(TableQuestions).
	Name("question_status")
