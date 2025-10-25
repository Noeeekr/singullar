package models

type QuestionListDifficulty int

const (
	Fundamental QuestionListDifficulty = iota
	Beginner
	Intermediare
	Advanced
	Expert
)

var EnumQuestionListDifficulty = EnvironmentDatabase.RegisterEnum(
	"question_list_difficulty",
	Fundamental,
	Beginner,
	Intermediare,
	Advanced,
	Expert,
)

var QuestionListDifficulties = map[QuestionListDifficulty]string{
	Fundamental:  "fundamental",
	Beginner:     "iniciante",
	Intermediare: "intermediario",
	Advanced:     "avançado",
	Expert:       "profissional",
}
