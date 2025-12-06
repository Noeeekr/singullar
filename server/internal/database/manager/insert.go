package manager

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/scan"
	"golang.org/x/crypto/bcrypt"
)

// Hash the password
func (ops *Operator) InsertManyUsers(requests ...*models.CreateUsers) ([]*models.Users, error) {
	var args []any = []any{}
	for _, request := range requests {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
		if err != nil {
			return nil, err
		}
		args = append(args, time.Now(), time.Now(), request.Name, request.Email, string(password), request.InstitutionId, request.Role, request.Segment)
	}

	var users []*models.Users
	query := models.TableUsers.
		Insert("created_at", "updated_at", "name", "email", "password", "institution_id", "role", "segment").
		Values(args...).
		Returning("created_at", "updated_at", "deleted_at", "name", "email", "password", "institution_id", "role", "id", "profile_picture", "segment").
		Scanner(scan.Users(&users))

	return users, ops.Do(query)
}

func (ops *Operator) InsertClassStudents(classId int, studentsIds *[]int) error {
	values := make([]any, len(*studentsIds)*2)
	{
		j := 0
		for i, id := range *studentsIds {
			i = i * 2
			j = i + 1
			values[i] = id
			values[j] = classId
		}
	}
	query := models.TableUsersClasses.
		Insert("user_id", "class_id").
		Values(values...)
	return ops.Do(query)
}
func (ops *Operator) InsertSubject(institutionId int, subject *models.CreateSubjectsRequest) (*models.Subjects, error) {
	subjects := []*models.Subjects{}
	query := models.TableSubjects.
		Select("id", "subject_name", "institution_id").
		Scanner(scan.Subjects(&subjects)).
		ThrowErrorOnFound()
	query.Where(query.And(
		query.Field("institution_id").IsEqual(institutionId),
		query.Field("subject_name").IsEqual(subject.SubjectName),
	))

	if err := ops.Do(query); err != nil {
		return nil, err
	}

	query = models.TableSubjects.
		Insert("subject_name", "institution_id").
		Values(subject.SubjectName, institutionId).
		Returning("id", "subject_name", "institution_id").
		Scanner(scan.Subjects(&subjects))
	if err := ops.Do(query); err != nil {
		return nil, err
	}
	return subjects[0], nil
}

// Throw borm.ErrNotFound when no questions is found for question list
func (ops *Operator) InsertQuestionList(institutionId int, questionList *models.CreateQuestionListRequest) (*models.QuestionLists, error) {
	targetIds := make([]any, len(questionList.QuestionIds))
	for i, id := range questionList.QuestionIds {
		targetIds[i] = id
	}

	ids := []int{}
	query := models.TableQuestions.
		Select("id").
		Scanner(scan.Integers(&ids))
	query.Where(query.And(
		query.Field("question_institution_id").IsEqual(institutionId),
		query.Field("id").IsAny(targetIds...),
	))

	// Query for valid question ids
	err := ops.Do(query)
	if err != nil {
		return nil, err
	}

	// If no valid ids, return not found
	if len(ids) == 0 {
		return nil, borm.ErrNotFound
	}

	lists := []*models.QuestionLists{}
	err = ops.Do(
		models.TableQuestionLists.
			Insert("created_at", "updated_at", "institution_id", "question_list_title", "subject_id", "question_list_difficulty_level").
			Values(time.Now(), time.Now(), institutionId, questionList.Title, questionList.SubjectId, questionList.DifficultyLevel).
			Returning("id", "created_at", "updated_at", "institution_id", "question_list_title", "subject_id", "question_list_difficulty_level").
			Scanner(scan.QuestionLists(&lists)),
	)
	if err != nil {
		return nil, err
	}

	values := make([]any, len(questionList.QuestionIds)*2)
	{
		offset := 0
		for i := range len(questionList.QuestionIds) {
			values[offset] = questionList.QuestionIds[i]
			values[offset+1] = lists[0].Id
			offset += 2
		}
	}
	ops.Do(
		models.TableQuestionListsQuestions.
			Insert("question_id", "question_list_id").
			Values(values...),
	)

	return lists[0], nil
}

func (ops *Operator) InsertQuestion(institutionId int, question *models.CreateQuestionRequest) (*models.Questions, error) {
	var questions []*models.Questions = []*models.Questions{}

	// Listen, I know questionDifficultyLevel needs to be checked agaisn't database to see if it as existing one and not an custom number, but wtv, it's not worth it.
	query := models.TableQuestions.
		Insert("question_institution_id", "created_at", "updated_at", "question_title", "question_description", "question_short_description", "question_difficulty_level", "question_correct_alternative", "question_subject_id").
		Values(institutionId, time.Now(), time.Now(), question.QuestionTitle, question.QuestionDescription, question.QuestionShortDescription, question.QuestionDifficultyLevel, question.QuestionCorrectAlternative, question.QuestionSubjectId).
		Returning("id", "question_institution_id", "created_at", "updated_at", "question_title", "question_description", "question_short_description", "question_difficulty_level", "question_correct_alternative", "question_subject_id").
		Scanner(scan.Questions(&questions))
	if err := ops.Do(query); err != nil {
		return nil, err
	}

	createdQuestion := questions[0]

	values := make([]any, len(question.Alternatives)*3)
	offset := 0
	for _, alternative := range question.Alternatives {
		values[offset] = createdQuestion.Id
		values[offset+1] = alternative
		values[offset+2] = question.QuestionCorrectAlternative == alternative
		offset += 3
	}

	query = models.TableQuestionAlternatives.
		Insert("question_id", "alternative", "is_correct").
		Values(values...)
	if err := ops.Do(query); err != nil {
		return nil, err
	}

	return createdQuestion, nil
}
func (ops *Operator) InsertClass(request *models.CreateClasses) (*models.Classes, error) {
	var classes []*models.Classes = []*models.Classes{}

	query := models.TableClasses.
		Insert("created_at", "updated_at", "name", "segment", "series", "institution_id", "teacher_id").
		Values(time.Now(), time.Now(), request.Name, request.Segment, request.Series, request.InstitutionId, request.TeacherId).
		Returning("created_at", "updated_at", "deleted_at", "name", "segment", "series", "institution_id", "id", "teacher_id").
		Scanner(scan.Classes(&classes))

	if err := ops.Do(query); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return classes[0], nil
}

// Creates the institutions and returns the administrator users
func (ops *Operator) InsertInstitutions(requests ...*models.CreateInstitutions) ([]*models.Users, error) {
	var args []any
	for _, request := range requests {
		args = append(args, time.Now(), time.Now(), request.Name)
	}

	institutionIds := []int{}
	query := models.TableInstitutions.
		Insert("created_at", "updated_at", "name").
		Values(args...).
		Returning("id").
		Scanner(scan.InstitutionsIds(&institutionIds))

	if err := ops.Do(query); err != nil {
		return nil, err
	}

	createUserRequests := make([]*models.CreateUsers, len(requests))
	for i, request := range requests {
		createUserRequests[i] = models.CreateUser("Administrator", request.Email, request.Password, institutionIds[0], models.ADMIN, nil)
	}

	return ops.InsertManyUsers(createUserRequests...)
}

// Requests are ordered ascending by issuerId and then title
func (ops *Operator) InsertNotifications(requests ...*NotificationRequest) ([]*models.NotificationContents, error) {
	sort.Slice(requests, func(i, j int) bool {
		leftIssuerID := requests[i].Content.IssuerId
		rightIssuerID := requests[j].Content.IssuerId

		// Sort by the smallest ID
		if leftIssuerID < rightIssuerID {
			return true
		}
		if leftIssuerID > rightIssuerID {
			return false
		}

		// Sort by the smallest title if they have the same ID
		result := strings.Compare(requests[j].Content.Title, requests[i].Content.Title)
		return result < 0
	})

	creationTime := time.Now()

	var notificationContentArgs []any = make([]any, len(requests)*5)
	for i, request := range requests {
		offset := 5 * i
		notificationContentArgs[offset] = creationTime
		notificationContentArgs[offset+1] = creationTime
		notificationContentArgs[offset+2] = request.Content.Title
		notificationContentArgs[offset+3] = request.Content.Description
		notificationContentArgs[offset+4] = request.Content.IssuerId
	}

	// Returns the notification in the same order that the requests are sorted
	var notifications []*models.NotificationContents
	query := models.TableNotificationContents.
		Insert("created_at", "updated_at", "title", "description", "issuer_id").
		Values(notificationContentArgs...).
		Returning("created_at", "updated_at", "deleted_at", "id", "issuer_id", "title", "description").
		Scanner(scan.NotificationContents(&notifications))
	if err := ops.Do(query); err != nil {
		return nil, err
	}

	sort.Slice(notifications, func(i, j int) bool {
		if notifications[i].IssuerId < notifications[j].IssuerId {
			return true
		}
		if notifications[i].IssuerId > notifications[j].IssuerId {
			return false
		}
		result := strings.Compare(notifications[i].Title, notifications[j].Title)
		return result < 0
	})

	userNotificationsArgs := []any{}
	for i, request := range requests {
		for _, user := range request.Users {
			userNotificationsArgs = append(
				userNotificationsArgs,
				user.UserId,
				user.UserRole,
				notifications[i].Id,
			)
		}
	}

	query = models.TableUsersNotifications.
		Insert("user_id", "user_role", "notification_id").
		Values(userNotificationsArgs...).
		Returning("user_id", "user_role", "notification_id")
	return notifications, ops.Do(query)
}
