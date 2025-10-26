package models

type Question struct {
	DefaultFields
	QuestionName            string `borm:"(NAME, question_name)"`
	QuestionDifficultyLevel int    `borm:"(NAME, question_difficulty_level)"`
	QuestionDescription     string `borm:"(NAME, question_description)"`
}

var TableQuestion = EnvironmentDatabase.
	RegisterTable(Question{}).
	NeedTables(TableQuestionDifficulty)
