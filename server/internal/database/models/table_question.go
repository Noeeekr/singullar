package models

// Request to get question data
type GetQuestionRequest struct {
	QuestionTitle           *string `borm:"(NAME, question_title)" json:"question_title,omitempty" binding:"required"`
	QuestionSubjectId       *int    `borm:"(NAME, question_subject_id) (FOREIGN KEY, subjects, id)" json:"question_subject_id,omitempty" binding:"required"`
	QuestionDifficultyLevel *int    `borm:"(NAME, question_difficulty_level) (FOREIGN KEY, question_difficulties, difficulty_level)" json:"question_difficulty_level,omitempty" binding:"required"`
}

// Request populated with necessary client fields to create a question
type CreateQuestionRequest struct {
	QuestionTitle              string   `borm:"(NAME, question_title)" json:"question_title,omitempty" binding:"required"`
	QuestionSubjectId          int      `borm:"(NAME, question_subject_id) (FOREIGN KEY, subjects, id)" json:"question_subject_id,omitempty" binding:"required"`
	QuestionDifficultyLevel    int      `borm:"(NAME, question_difficulty_level) (FOREIGN KEY, question_difficulties, difficulty_level)" json:"question_difficulty_level,omitempty" binding:"required"`
	QuestionDescription        string   `borm:"(NAME, question_description)" json:"question_description" binding:"required"`
	QuestionShortDescription   string   `borm:"(NAME, question_short_description)" json:"question_short_description" binding:"required"`
	QuestionCorrectAlternative string   `borm:"(NAME, question_correct_alternative)" json:"question_correct_alternative" binding:"required"`
	Alternatives               []string `borm:"(IGNORE)" json:"alternatives" binding:"required"`
}

// Request populated with server sensitive information to create a question
type CreateQuestions struct {
	DefaultFields
	CreateQuestionRequest
	QuestionInstitutionId int `borm:"(NAME, question_institution_id) (FOREING KEY, institutions, id)"`
}

// Request after being populated by database
type Questions struct {
	ID
	CreateQuestions
}

// Target: Client + User Specific Information
// Comes From: Database return + Extra Query
type ExtendedQuestions struct {
	Questions

	// Not implemented yet
	// Incomplete | Completed | Untouched
	Status int `borm:"(FOREIGN KEY, question_status, id)" json:"question_status"`
}

var TableQuestions = EnvironmentDatabase.
	RegisterTable(Questions{}).
	NeedTables(TableQuestionDifficulty, TableSubjects)
