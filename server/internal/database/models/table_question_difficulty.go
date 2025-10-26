package models

type QuestionDifficulty struct {
	DifficultyLevel int `borm:"(NAME, difficulty_level)"`
	DifficultyName  int `borm:"(NAME, difficulty_name)"`
}

var TableQuestionDifficulty = EnvironmentDatabase.
	RegisterTable(QuestionDifficulty{}).
	Name("question_difficulty")
