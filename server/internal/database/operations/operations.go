// package operations provides reliable database operations for the api
package operations

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

// Operations organizes and separates database operation logic. For asyncronous implementations a new instance of Operations must be created for each goroutine.
type Operations struct {
	currentTransction *borm.Transaction
	*borm.Commiter
}

// Returns an instance of Operations. It is recommended to check [type Operations] for further instructions on how to use it.
func New(commiter *borm.Commiter) *Operations {
	return &Operations{
		currentTransction: nil,
		Commiter:          commiter,
	}
}

// StartTransaction starts a new transaction that lasts until the next CommitTransaction()
func (ops *Operations) StartTransaction() error {
	tx, err := ops.Commiter.StartTx()
	ops.currentTransction = tx
	return err
}

// Commits any ongoing transaction
func (ops *Operations) CommitTransaction() error {
	return ops.currentTransction.Commit()
}

// Hash the password
func (ops *Operations) InsertManyUsers(requests ...*models.CreateUsers) ([]*models.Users, error) {
	var args []any = []any{}
	for _, request := range requests {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
		if err != nil {
			return nil, err
		}
		args = append(args, time.Now(), time.Now(), request.Name, request.Email, string(password), request.InstitutionId, request.Role, request.Segment)
	}

	// InsertMany: transactions.NewRequest(fmt.Sprintf(`
	// 		INSERT INTO %s (created_at, updated_at, name, email, password, institution_id, role, segment)
	// 		VALUES %s
	// 		RETURNING created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture, segment;
	// 	`, usersTableName, placeholder)).AllowValueRepeat(placeholder, 8),
	var users []*models.Users
	err := ops.currentTransction.Do(models.TableUsers.
		Insert("created_at", "updated_at", "name", "email", "password", "institution_id", "role", "segment").
		Values(args...).
		Returning("created_at", "updated_at", "deleted_at", "name", "email", "password", "institution_id", "role", "id", "profile_picture", "segment").
		Scanner(scan.Users(&users)),
	)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ops *Operations) InsertClassStudents(classId int, studentsIds *[]int) error {
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
	return ops.currentTransction.Do(
		models.TableUsersClasses.
			Insert("user_id", "class_id").
			Values(values...),
	)
}
func (ops *Operations) InsertSubject(institutionId int, subject *models.CreateSubjectsRequest) (*models.Subjects, error) {
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

	if err := ops.StartTransaction(); err != nil {
		return nil, err
	}

	query = models.TableSubjects.
		Insert("subject_name", "institution_id").
		Values(subject.SubjectName, institutionId).
		Returning("id", "subject_name", "institution_id").
		Scanner(scan.Subjects(&subjects))
	if err := ops.currentTransction.Do(query); err != nil {
		return nil, err
	}
	return subjects[0], ops.currentTransction.Commit()
}
func (ops *Operations) InsertQuestion(institutionId int, question *models.CreateQuestionRequest) (*models.Questions, error) {
	var questions []*models.Questions = []*models.Questions{}
	err := ops.StartTransaction()
	if err != nil {
		return nil, err
	}

	err = ops.currentTransction.Do(
		models.TableQuestions.
			Insert("question_institution_id", "created_at", "updated_at", "question_title", "question_description", "question_short_description", "question_difficulty_level", "question_correct_alternative", "question_subject_id").
			// Listen, I know questionDifficultyLevel needs to be checked agaisn't database to see if it as existing one and not an custom number, but wtv, it's not worth it.
			Values(institutionId, time.Now(), time.Now(), question.QuestionTitle, question.QuestionDescription, question.QuestionShortDescription, question.QuestionDifficultyLevel, question.QuestionCorrectAlternative, question.QuestionSubjectId).
			Returning("id", "question_institution_id", "created_at", "updated_at", "question_title", "question_description", "question_short_description", "question_difficulty_level", "question_correct_alternative", "question_subject_id").
			Scanner(scan.Questions(&questions)),
	)

	if err != nil {
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
	err = ops.currentTransction.Do(
		models.TableQuestionAlternatives.
			Insert("question_id", "alternative", "is_correct").
			Values(values...),
	)
	if err != nil {
		return nil, err
	}

	return createdQuestion, ops.currentTransction.Commit()
}
func (ops *Operations) InsertClass(request *models.CreateClasses) (*models.Classes, error) {
	var classes []*models.Classes = []*models.Classes{}
	err := ops.currentTransction.Do(
		models.TableClasses.
			Insert("created_at", "updated_at", "name", "segment", "series", "institution_id", "teacher_id").
			Values(time.Now(), time.Now(), request.Name, request.Segment, request.Series, request.InstitutionId, request.TeacherId).
			Returning("created_at", "updated_at", "deleted_at", "name", "segment", "series", "institution_id", "id", "teacher_id").
			Scanner(scan.Classes(&classes)),
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return classes[0], nil
}

// Creates the institutions and returns the administrator users
func (ops *Operations) InsertInstitutions(requests ...*models.CreateInstitutions) ([]*models.Users, error) {
	var args []any
	for _, request := range requests {
		args = append(args, time.Now(), time.Now(), request.Name)
	}

	var ids []int
	err := ops.currentTransction.Do(models.TableInstitutions.
		Insert("created_at", "updated_at", "name").
		Values(args...).
		Returning("id").
		Scanner(scan.InstitutionsIds(&ids)),
	)
	if err != nil {
		return nil, err
	}

	createUserRequests := make([]*models.CreateUsers, len(requests))
	for i, request := range requests {
		createUserRequests[i] = models.CreateUser("Administrator", request.Email, request.Password, ids[0], models.ADMIN, nil)
	}

	return ops.InsertManyUsers(createUserRequests...)
}

// Requests are ordered ascending by issuerId and then title
func (ops *Operations) InsertNotifications(requests ...*NotificationRequest) ([]*models.NotificationContents, error) {
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
	err := ops.currentTransction.Do(
		models.TableNotificationContents.
			Insert("created_at", "updated_at", "title", "description", "issuer_id").
			Values(notificationContentArgs...).
			Returning("created_at", "updated_at", "deleted_at", "id", "issuer_id", "title", "description").
			Scanner(scan.NotificationContents(&notifications)),
	)
	if err != nil {
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

	err = ops.currentTransction.Do(
		models.TableUsersNotifications.
			Insert("user_id", "user_role", "notification_id").
			Values(userNotificationsArgs...).
			Returning("user_id", "user_role", "notification_id"),
	)
	return notifications, err
}
func (ops *Operations) DeleteUserById(id int) error {
	query := models.TableUsers.Delete()
	query.Where(query.Field("id").IsEqual(id))
	err := ops.currentTransction.Do(query)
	return err
}
func (ops *Operations) DeleteUserByEmail(email string) error {
	query := models.TableUsers.Delete()
	query.Where(query.Field("email").IsEqual(email))
	err := ops.currentTransction.Do(query)
	return err
}
func (ops *Operations) DeleteInstitutionByName(name string) error {
	var ids []int
	query := models.TableInstitutions.Delete()
	query.Where(query.Field("name").IsEqual(name)).
		Returning("id").
		Scanner(scan.Integers(&ids))
	err := ops.currentTransction.Do(query)

	return err
}
func (ops *Operations) DeleteInstitutionById(id int) error {
	query := models.TableInstitutions.Delete()
	query.Where(query.Field("id").IsEqual(id))
	err := ops.currentTransction.Do(query)
	return err
}

func (ops *Operations) DeleteNotificationsByIssuerId(ids ...int) error {
	for _, id := range ids {
		query := models.TableNotificationContents.Delete()
		query.Where(query.Field("id").IsEqual(id))
		err := ops.currentTransction.Do(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func addSelectClassesFilters(query *borm.Query, filters *[]*FilterClassesOptions) *borm.ConditionalQuery {
	conditions := make([]*borm.ConditionalQuery, len(*filters))
	for i, filter := range *filters {
		condition := []*borm.ConditionalQuery{}
		if filter.Id != nil {
			condition = append(condition, query.Field("c.id").IsEqual(*filter.Id))
		}
		if filter.ClassName == nil {
			condition = append(condition, query.Field("c.name").IsLike("%", false))
		} else {
			condition = append(condition, query.Field("c.name").IsLike("%"+*filter.ClassName+"%", false))
		}
		if filter.StudentName != nil {
			condition = append(condition, query.Compose(query.And(
				query.Field("u.name").IsLike("%"+strings.ToLower(*filter.StudentName)+"%", false),
				query.Field("u.role").IsEqual(models.STUDENT),
			)))
		}
		if filter.TeacherName != nil {
			condition = append(condition, query.Compose(query.And(
				query.Field("u.name").IsLike("%"+strings.ToLower(*filter.TeacherName)+"%", false),
				query.Field("u.role").IsEqual(models.TEACHER),
			)))
		}
		if filter.Series != nil {
			condition = append(condition, query.Field("c.series").IsEqual(*filter.Series))
		}
		if filter.Segment != nil {
			condition = append(condition, query.Field("c.segment").IsEqual(*filter.Segment))
		}
		if filter.CreationYear != nil {
			// query.And("c.created_at").After(filter.CreationYear) => Which is time.Time()
		}
		conditions[i] = query.Compose(query.And(condition...))
	}
	return query.Compose(query.Or(conditions...))
}
func addSelectStudentsFilters(query *borm.Query, filters *[]FilterStudentOptions) *borm.ConditionalQuery {
	conditions := make([]*borm.ConditionalQuery, len(*filters))
	for i, filter := range *filters {
		condition := []*borm.ConditionalQuery{}
		if filter.Segment == models.UNKNOWN_SEGMENT {
			condition = append(condition, query.Field("u.segment").IsEqual(nil))
		} else {
			condition = append(condition, query.Field("u.segment").IsEqual(filter.Segment))
		}
		if filter.Email != "" {
			condition = append(condition, query.Field("u.email").IsEqual(filter.Email))
		}
		if filter.Name != "" {
			condition = append(condition, query.Field("u.name").IsLike("%"+strings.ToLower(filter.Name)+"%", false))
		}
		if filter.ID != nil {
			condition = append(condition, query.Field("u.id").IsEqual(filter.ID))
		}
		if filter.ClassID != nil {
			condition = append(condition, query.Field("c.id").IsEqual(filter.ClassID))
		}
		conditions[i] = query.And(condition...)
	}
	return query.Compose(query.Or(conditions...))
}
