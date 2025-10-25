package models

// Stage: Pre server field population (Request to API)
type CreateQuestionListRequest struct {
	Title           string `json:"title" binding:"required"`
	Subject         int    `borm:"(FOREIGN KEY, subjects, id)" json:"subject" binding:"required"`
	DifficultyLevel int    `borm:"(NAME, difficulty_level)" json:"difficulty_level" binding:"required"`
}

// Stage: Post Server field population
type CreateQuestionList struct {
	CreateQuestionListRequest

	// Defaults to zero since there are no questions added yet
	QuestionQuantity int `borm:"(NAME, question_quantity) (DEFAULT, 0)" json:"question_quantity" binding:"required"`
	// Infered based on user information
	InstitutionId int `borm:"(NAME, institution_id) (FOREIGN KEY, institutions, id)" json:"institution_id" binding:"required"`
	// Infered based on Difficulty Level
	DifficultyName string `borm:"(NAME, difficulty_name)" json:"difficulty_name" binding:"required"`
}

// Stage: Post Database field population (Database return)
type QuestionList struct {
	DefaultFields
	CreateQuestionList
}

var TableQuestionList = EnvironmentDatabase.
	RegisterTable(QuestionList{}).
	NeedTables(TableQuestion).
	NeedRoles(EnumQuestionListDifficulty)
