package models

type QuestionDifficulty struct {
	DifficultyLevel int    `borm:"(NAME, difficulty_level) (CONSTRAINTS, PRIMARY KEY)" json:"question_difficulty_level"`
	DifficultyName  string `borm:"(NAME, difficulty_name) (CONSTRAINTS, UNIQUE, NOT NULL)" json:"question_difficulty_name"`
}

var TableQuestionDifficulty = EnvironmentDatabase.
	RegisterTable(QuestionDifficulty{}).
	Name("question_difficulties")
