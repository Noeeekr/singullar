package operations

import (
	"errors"
	"time"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/scan"
)

type FilterClassesOptions struct {
	Id           *int                 `json:"id"`
	ClassName    *string              `json:"class_name"`
	TeacherName  *string              `json:"teacher_name"`
	StudentName  *string              `json:"student_name"`
	Series       *string              `json:"series"`
	Segment      *models.UserSegments `json:"segment"`
	CreationYear *time.Time           `json:"creation_year"`
}

func (ops *Operations) SelectClasses(institutionId int, filters ...*FilterClassesOptions) (*[]*models.Classes, error) {
	classes := &[]*models.Classes{}

	query := models.TableClasses.
		SelectDistinct("c.created_at", "c.updated_at", "c.deleted_at", "c.name", "c.segment", "c.series", "c.institution_id", "c.id", "c.teacher_id").As("c").
		LeftJoin(models.TableUsersClasses, "uc").On("c.id", "uc.class_id").
		LeftJoin(models.TableUsers, "u").On("uc.user_id", "u.id")
	condition := query.Field("c.institution_id").IsEqual(institutionId)
	if len(filters) != 0 {
		condition = query.And(condition, addSelectClassesFilters(query, &filters))
	}
	query.Where(condition)

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
	query.Where(query.Field("u.email").IsAny(emailList...))
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
			query.Field("u.id").IsAny(idList...),
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

type SelectUsersOptions struct {
	Roles   []*models.UserRole          `json:"accepted_roles" binding:"omitempty"`
	Filters []*SelectUsersFilterOptions `json:"filters" binding:"omitempty"`
	Offset  int                         `json:"offset"`
}
type SelectUsersFilterOptions struct {
	ID *int `json:"id" binding:"omitempty"`
}

// Returns an empty array if no users were found. Hashed Password is returned and must be removed
func (ops *Operations) SelectUsers(institutionId int, request *SelectUsersOptions) ([]*models.Users, error) {
	users := []*models.Users{}
	query := models.TableUsers.
		Select("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
		InnerJoin(models.TableInstitutions, "i").On("i.id", "u.institution_id").
		Scanner(scan.Users(&users))
	query.Where(query.And(
		query.Field("u.institution_id").IsEqual(institutionId),
		addSelectUsersFilters(query, request),
	)).Limit(10).Offset(request.Offset)
	//$$$
	err := ops.Do(query)
	if err != nil && !errors.Is(err, borm.ErrNotFound) {
		return nil, err
	}
	return users, nil
}

func addSelectUsersFilters(query *borm.Query, request *SelectUsersOptions) *borm.ConditionalQuery {
	conditions := []*borm.ConditionalQuery{}
	if len(request.Roles) != 0 {
		targetRoles := make([]any, len(request.Roles))
		for i, role := range request.Roles {
			targetRoles[i] = role
		}
		conditions = append(conditions, query.Field("u.role").IsAny(targetRoles...))
	}

	filters := []*borm.ConditionalQuery{}
	for _, filter := range request.Filters {
		var condition []*borm.ConditionalQuery = []*borm.ConditionalQuery{}
		if filter.ID != nil {
			condition = append(condition, query.Field("u.id").IsEqual(*filter.ID))
		}
		if len(condition) > 0 {
			filters = append(filters, query.Compose(query.And(condition...)))
		}
	}

	conditions = append(conditions, query.Compose(query.And(filters...)))
	return query.And(conditions...)
}

type FilterStudentOptions struct {
	Email   string              `json:"email" binding:"omitempty,email"`
	Name    string              `json:"name" binding:"omitempty,min=1"`
	Segment models.UserSegments `json:"segment" binding:"omitempty"`
	ID      *int                `json:"id" binding:"omitempty"`
	ClassID *int                `json:"class_id" binding:"omitempty"`
}

// Returns an empty array if no users were found. Hashed Password is returned and must be removed. Default values will be ignored in search, expect for segment.
func (ops *Operations) SelectStudents(institutionId int, options *[]FilterStudentOptions, includePassword bool) ([]*models.Users, error) {
	users := []*models.Users{}
	query := models.TableUsers.
		SelectDistinct("u.created_at", "u.updated_at", "u.deleted_at", "u.name", "u.email", "u.password", "u.institution_id", "u.role", "u.id", "u.profile_picture", "u.segment").As("u").
		InnerJoin(models.TableInstitutions, "i").On("i.id", "u.institution_id").
		LeftJoin(models.TableUsersClasses, "uc").On("u.id", "uc.user_id").
		LeftJoin(models.TableClasses, "c").On("c.id", "uc.class_id").
		Scanner(scan.Users(&users))
	query.Where(
		query.And(
			query.Field("u.institution_id").IsEqual(institutionId),
			query.Field("u.role").IsEqual(models.STUDENT),
			addSelectStudentsFilters(query, options),
		),
	)

	err := ops.Do(query)
	if err != nil && !errors.Is(err, borm.ErrNotFound) {
		return nil, err
	}

	if !includePassword {
		for i := range users {
			users[i].Password = ""
		}
	}
	return users, nil
}

func (ops *Operations) SelectSubjects(InstitutionId int) (*[]*models.Subjects, error) {
	subjects := []*models.Subjects{}
	query := models.TableSubjects.
		Select("s.id, s.subject_name", "s.institution_id").As("s").
		InnerJoin(models.TableInstitutions, "i").On("i.id", "s.institution_id").
		Scanner(scan.Subjects(&subjects))
	query.Where(query.Field("s.institution_id").IsEqual(InstitutionId))
	if err := ops.Commiter.Do(query); err != nil {
		return nil, err
	}
	return &subjects, nil
}
