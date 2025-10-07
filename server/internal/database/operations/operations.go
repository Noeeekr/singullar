// package operations provides reliable database operations for the api
package operations

import (
	"errors"
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

type FilterClassOptions struct {
	ClassName    *string              `json:"class_name"`
	TeacherName  *string              `json:"teacher_name"`
	StudentName  *string              `json:"student_name"`
	Series       *string              `json:"series"`
	Segment      *models.UserSegments `json:"segment"`
	CreationYear *time.Time           `json:"creation_year"`
}

func (ops *Operations) SelectClasses(institutionId int, filters ...*FilterClassOptions) (*[]*models.Classes, error) {
	classes := &[]*models.Classes{}

	query := models.TableClasses.
		Select("c.created_at", "c.updated_at", "c.deleted_at", "c.name", "c.segment", "c.series", "c.institution_id", "c.id", "c.teacher_id").As("c").
		LeftJoin(models.TableUsersClasses, "uc").On("c.id", "uc.class_id").
		LeftJoin(models.TableUsers, "u").On("uc.user_id", "u.id")
	query.Where(
		query.And(
			query.Field("c.institution_id").IsEqual(institutionId),
			addSelectClassesFilters(query, &filters),
		),
	)

	err := ops.Commiter.Do(query.Scanner(scan.Classes(classes)))
	if err != nil {
		return nil, err
	}
	return classes, nil
}
func (ops *Operations) SelectNotificationsByTargetId(id int) (*[]*models.Notifications, error) {
	var notifications []*models.Notifications
	query := models.TableNotificationContents.
		Select("n.created_at", "n.updated_at", "n.deleted_at", "u.id", "u.name", "n.id", "n.title", "n.description").As("n").
		InnerJoin(models.TableUsersNotifications, "un").On("un.notification_id", "n.id").
		InnerJoin(models.TableUsers, "u").On("u.id", "un.user_id").
		Scanner(scan.Notifications(&notifications))
	query.Where(query.Field("u.id").IsEqual(id))
	err := ops.Commiter.Do(query)
	return &notifications, err
}

// Returns ErrNotFound, ErrSyntax, ErrFailedTransaction
func (ops *Operations) SelectUserByEmail(emails ...string) ([]*models.Users, error) {
	var users []*models.Users
	var emailList []any = make([]any, len(emails))
	for i, v := range emails {
		emailList[i] = v
	}
	query := models.TableUsers.
		Select("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
		Scanner(scan.Users(&users))
	query.Where(query.Field("u.email").IsIn(emailList...))
	err := ops.Commiter.Do(query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ops *Operations) SelectUsersById(institutionId int, ids ...int) ([]*models.Users, error) {
	var users []*models.Users
	var idList []any = make([]any, len(ids))
	for i, v := range ids {
		idList[i] = v
	}
	query := models.TableUsers.
		Select("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
		Scanner(scan.Users(&users))
	query.Where(
		query.And(
			query.Field("u.institution_id").IsEqual(institutionId),
			query.Field("u.id").IsIn(idList...),
		),
	)
	err := ops.Do(query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ops *Operations) SelectInstitutionByName(name string) (*models.Institutions, error) {
	var insts []*models.Institutions
	query := models.TableInstitutions.
		Select("i.created_at", "i.updated_at", "i.deleted_at", "i.name", "i.id").As("i").
		Scanner(scan.Institutions(&insts))
	query.Where(query.Field("name").IsEqual(name))
	err := ops.Do(query)
	if err != nil {
		return nil, err
	}
	return insts[0], nil
}
func (ops *Operations) SelectInstitutionById(id int) (*models.Institutions, error) {
	var insts []*models.Institutions = []*models.Institutions{}
	query := models.TableInstitutions.
		Select("i.created_at", "i.updated_at", "i.deleted_at", "i.name", "i.id").As("i").
		Scanner(scan.Institutions(&insts))
	query.Where(query.Field("id").IsEqual(id))
	err := ops.Do(query)
	if err != nil {
		return nil, err
	}
	return insts[0], nil
}

// Returns an empty array if no users were found. Hashed Password is returned and must be removed
func (ops *Operations) SelectUsers(institutionId int, roles ...models.UserRole) ([]*models.Users, error) {
	if len(roles) == 0 {
		roles = []models.UserRole{models.STUDENT}
	}

	var targetRoles = make([]any, len(roles))
	for i, role := range roles {
		targetRoles[i] = role
	}

	users := []*models.Users{}
	query := models.TableUsers.
		Select("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
		InnerJoin(models.TableInstitutions, "i").On("i.id", "u.institution_id").
		Scanner(scan.Users(&users))
	query.
		Where(
			query.And(
				query.Field("u.institution_id").IsEqual(institutionId),
				query.Field("u.role").IsIn(targetRoles...)))
	err := ops.Do(query)
	if err != nil && !errors.Is(err, borm.ErrNotFound) {
		return nil, err
	}
	return users, nil
}

type FilterStudentOptions struct {
	Email   string              `json:"email" binding:"omitempty,email"`
	Name    string              `json:"name" binding:"omitempty,min=1"`
	Segment models.UserSegments `json:"segment" binding:"omitempty"`
	ID      string              `json:"id" binding:"omitempty"`
}

// Returns an empty array if no users were found. Hashed Password is returned and must be removed. Default values will be ignored in search, expect for segment.
func (ops *Operations) SelectStudents(institutionId int, options *[]FilterStudentOptions, includePassword bool) ([]*models.Users, error) {
	users := []*models.Users{}
	query := models.TableUsers.
		Select("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
		InnerJoin(models.TableInstitutions, "i").On("i.id", "u.institution_id").
		Scanner(scan.Users(&users))
	query.Where(
		query.And(
			query.Field("u.institution_id").IsEqual(institutionId),
			query.Field("u.role").IsEqual(models.STUDENT),
			addSelectStudentsFilters(query, options),
		),
	)

	blocks := make([]string, len(query.Blocks))
	fmt.Println(len(*options))
	for i := range query.Blocks {
		blocks[i] = query.Blocks[i].Block
	}
	fmt.Printf("[%s]\n", strings.Join(blocks, "]\n["))
	err := ops.Do(query)
	if err != nil && !errors.Is(err, borm.ErrNotFound) {
		return nil, err
	}

	if !includePassword {
		for i := range users {
			fmt.Println(users[i].Name, users[i].Segment, users[i].Id)
			users[i].Password = ""
		}
	}
	return users, nil
}

/*er
student by name
student by segment - or without
student by email
student by id
*/

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

func addSelectClassesFilters(query *borm.Query, filters *[]*FilterClassOptions) *borm.ConditionalQuery {
	conditions := make([]*borm.ConditionalQuery, len(*filters))
	for i, filter := range *filters {
		condition := []*borm.ConditionalQuery{}
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

	borm.Settings().Environment().SetEnvironment(borm.DEBUGGING)
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
		if filter.ID != "" {
			condition = append(condition, query.Field("u.id").IsEqual(filter.ID))
		}
		conditions[i] = query.And(condition...)
	}
	return query.Compose(query.Or(conditions...))
}
