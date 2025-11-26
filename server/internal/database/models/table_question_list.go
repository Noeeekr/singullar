package models

type QuestionListClientFields struct {
	Title           string `borm:"(NAME, question_list_title)" json:"question_list_title" binding:"required"`
	DifficultyLevel int    `borm:"(NAME, question_list_difficulty_level) (FOREIGN KEY, question_difficulties, difficulty_level)" json:"question_list_difficulty_level" binding:"required"`
	SubjectId       int    `borm:"(NAME, subject_id) (FOREIGN KEY, subjects, id)" json:"question_list_subject_id" binding:"required"`
}

// Target: Necessary client information to server validation and field population.
// Comes from: Client
type CreateQuestionListRequest struct {
	QuestionListClientFields

	QuestionIds []int `json:"question_list_question_ids" binding:"required"`
}

// Target: Insertion into database
// Comes from: Client request after server field population
type CreateQuestionList struct {
	QuestionListClientFields

	// Infered based on user information
	InstitutionId int `borm:"(NAME, institution_id) (FOREIGN KEY, institutions, id)" json:"institution_id" binding:"required"`
}

// Target: Client Ready Data
// Comes From: Database return
type QuestionLists struct {
	ID
	DefaultFields
	CreateQuestionList
}

type QuestionListsFilters struct {
	Fields []struct {
		Title           *string `json:"title"`
		DifficultyLevel *int    `json:"difficulty_level"`
		SubjectId       *int    `json:"subject_id"`
	} `json:"fields"`
	Ids []int `json:"ids"`
}

// Target: Client Ready Data + Extra info
// Comes From: Database return
type ExtendedQuestionList struct {
	QuestionLists

	SubjectName      string `borm:"(IGNORE)" json:"subject_name" binding:"required"`
	QuestionQuantity int    `borm:"(IGNORE)" json:"question_quantity" binding:"required"`
}

var TableQuestionLists = EnvironmentDatabase.
	RegisterTable(QuestionLists{}).
	Name("question_lists").
	NeedTables(
		TableQuestionDifficulty,
		TableQuestions,
		TableInstitutions,
		TableSubjects,
	)
