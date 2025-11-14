package models

// Request sent from client
type CreateQuestionRequest struct {
	QuestionTitle            string `borm:"(NAME, question_title)" json:"question_title" binding:"required"`
	QuestionDescription      string `borm:"(NAME, question_description)" json:"question_description" binding:"required"`
	QuestionShortDescription string `borm:"(NAME, question_short_description)" json:"question_short_description" binding:"required"`
	QuestionDifficultyLevel  int    `borm:"(NAME, question_difficulty_level) (FOREIGN KEY, question_difficulties, difficulty_level)" json:"question_difficulty_level" binding:"required"`

	QuestionCorrectAlternative string   `borm:"(NAME, question_correct_alternative)" json:"question_correct_alternative" binding:"required"`
	Alternatives               []string `borm:"(IGNORE)" json:"alternatives" binding:"required"`
}

// Client request after populated by server
type CreateQuestions struct {
	DefaultFields
	CreateQuestionRequest
	QuestionInstitutionId int `borm:"(NAME, question_institution_id) (FOREING KEY, institutions, id)"`
}

// Client request after populated by server and database
type Questions struct {
	ID
	CreateQuestions
}

var TableQuestions = EnvironmentDatabase.
	RegisterTable(Questions{}).
	NeedTables(TableQuestionDifficulty)
