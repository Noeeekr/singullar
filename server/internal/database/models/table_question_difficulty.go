package models

type QuestionDifficulty struct {
	DifficultyLevel uint `borm:"(NAME, difficulty_level)"`
	DifficultyName  uint `borm:"(NAME, difficulty_name)"`
}

var TableQuestionDifficulty = EnvironmentDatabase.
	RegisterTable(QuestionDifficulty{}).
	Name("question_difficulty")
