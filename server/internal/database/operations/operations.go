// package operations provides reliable database operations for the api
package operations

import (
	"errors"
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

func (ops *Operations) SelectNotificationsByTargetId(id int) (*[]*models.Notifications, error) {
	var notifications []*models.Notifications
	err := ops.Commiter.Do(models.TableNotificationContents.
		Select("n.created_at", "n.updated_at", "n.deleted_at", "u.id", "u.name", "n.id", "n.title", "n.description").As("n").
		InnerJoin(models.TableUsersNotifications, "un").On("un.notification_id", "n.id").
		InnerJoin(models.TableUsers, "u").On("u.id", "un.user_id").
		Where("u.id").Equals(id).Scanner(scan.Notifications(&notifications)),
	)
	return &notifications, err
}

// Returns ErrNotFound, ErrSyntax, ErrFailedTransaction
func (ops *Operations) SelectUserByEmail(email string) (*models.Users, error) {
	var users []*models.Users
	err := ops.Commiter.Do(
		models.TableUsers.
			Select("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
			Where("u.email").Equals(email).
			Scanner(scan.Users(&users)),
	)
	if err != nil {
		return nil, err
	}
	return users[0], nil
}

func (ops *Operations) SelectUserById(id int) (*models.Users, error) {
	var users []*models.Users

	err := ops.Do(
		models.TableUsers.
			Select("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
			Where("u.id").Equals(id).
			Scanner(scan.Users(&users)),
	)
	if err != nil {
		return nil, err
	}

	return users[0], nil
}

func (ops *Operations) SelectInstitutionByName(name string) (*models.Institutions, error) {
	var insts []*models.Institutions
	err := ops.Do(
		models.TableInstitutions.
			Select("i.created_at", "i.updated_at", "i.deleted_at", "i.name", "i.id").As("i").
			Where("name").Equals(name).
			Scanner(scan.Institutions(&insts)),
	)
	if err != nil {
		return nil, err
	}
	return insts[0], nil
}
func (ops *Operations) SelectInstitutionById(id int) (*models.Institutions, error) {
	var insts []*models.Institutions = []*models.Institutions{}
	err := ops.Do(
		models.TableInstitutions.
			Select("i.created_at", "i.updated_at", "i.deleted_at", "i.name", "i.id").As("i").
			Where("id").Equals(id).
			Scanner(scan.Institutions(&insts)),
	)
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
	err := ops.Do(models.TableUsers.
		Select("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
		InnerJoin(models.TableInstitutions, "i").On("i.id", "u.institution_id").
		Where("u.institution_id").Equals(institutionId).
		Where("u.role").In(targetRoles...).
		Scanner(scan.Users(&users)),
	)
	if err != nil && !errors.Is(err, borm.ErrNotFound) {
		return nil, err
	}
	return users, nil
}

type SelectStudentsOptions struct {
	Email   string               `json:"email" binding:"omitempty,email"`
	Name    string               `json:"name" binding:"omitempty,min=1"`
	Segment *models.UserSegments `json:"segment" binding:"omitempty"`
	ID      *int                 `json:"id" binding:"omitempty"`
}

// Returns an empty array if no users were found. Hashed Password is returned and must be removed. Default values will be ignored in search, expect for segment.
func (ops *Operations) SelectStudents(institutionId int, options *SelectStudentsOptions) ([]*models.Users, error) {
	users := []*models.Users{}
	query := models.TableUsers.
		Select("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
		InnerJoin(models.TableInstitutions, "i").On("i.id", "u.institution_id").
		Scanner(scan.Users(&users)).
		Where("u.institution_id").Equals(institutionId).
		Where("u.role").Equals(models.STUDENT)
	if options != nil {
		if options.Segment != nil {
			query.Where("u.segment").Equals(nil)
		}
		if options.Email != "" {
			query.Where("u.email").Equals(options.Email)
		}
		if options.Name != "" {
			query.Where("u.name").Like("%"+strings.ToLower(options.Name)+"%", false)
		}
		if options.ID != nil {
			query.Where("u.id").Equals(*options.ID)
		}
	}

	err := ops.Do(query)
	if err != nil && !errors.Is(err, borm.ErrNotFound) {
		return nil, err
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
		if result >= 0 {
			return false
		}
		return true
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
		if result >= 0 {
			return false
		}
		return true
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
	err := ops.currentTransction.Do(
		models.TableUsers.
			Delete().
			Where("id").Equals(id),
	)
	return err
}
func (ops *Operations) DeleteUserByEmail(email string) error {
	err := ops.currentTransction.Do(
		models.TableUsers.
			Delete().
			Where("email").Equals(email),
	)
	return err
}
func (ops *Operations) DeleteInstitutionByName(name string) error {
	var ids []int
	err := ops.currentTransction.Do(
		models.TableInstitutions.
			Delete().
			Where("name").Equals(name).
			Returning("id").
			Scanner(scan.Integers(&ids)),
	)

	return err
}
func (ops *Operations) DeleteInstitutionById(id int) error {
	err := ops.currentTransction.Do(
		models.TableInstitutions.
			Delete().
			Where("id").Equals(id),
	)
	return err
}

func (ops *Operations) DeleteNotificationsByIssuerId(ids ...int) error {
	for _, id := range ids {
		err := ops.currentTransction.Do(models.TableNotificationContents.
			Delete().
			Where("id").Equals(id),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
