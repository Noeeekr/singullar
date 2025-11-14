package models

type QuestionAlternatives struct {
	QuestionID  int `borm:"(NAME, question_id) (FOREIGN KEY, questions, id)"`
	Alternative string
	IsCorrect   bool `borm:"(NAME, is_correct)"`
}

var TableQuestionAlternatives = EnvironmentDatabase.
	RegisterTable(QuestionAlternatives{}).
	Name("question_alternatives").
	NeedTables(TableQuestions)
