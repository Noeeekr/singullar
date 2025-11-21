package models

// Question List - Questions [Many-Many relation]
type QuestionListsQuestions struct {
	QuestionId     int `borm:"(NAME, question_id) (FOREIGN KEY, questions, id)"`
	QuestionListId int `borm:"(NAME, question_list_id) (FOREIGN KEY, question_lists, id)"`
}

var TableQuestionListsQuestions = EnvironmentDatabase.
	RegisterTable(QuestionListsQuestions{}).
	Name("question_lists_questions").
	NeedTables(TableQuestions, TableQuestionLists)
