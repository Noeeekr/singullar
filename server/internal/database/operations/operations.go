// package operations provides reliable database operations for the api
package operations

import (
	"database/sql"
	"time"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/scan"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type Operations struct {
	*transactions.Manager
}

// Returns an instance of Operations. Operations contains methods that to make the most used transactions instantly.
func New(db *sql.DB) *Operations {
	return &Operations{
		Manager: transactions.NewManager(db),
	}
}

func (ops *Operations) SelectUserByEmail(email string) (user *models.Users, res *common.Response) {
	var users []*models.Users

	res = ops.Query(
		models.UsersTable.Requests.SelectOneByEmail.
			WithArgs(email).
			WithRowsScanner(scan.Users(&users)),
	)
	if res != nil {
		return user, res
	}

	if len(users) > 1 {
		return nil, common.NewResponse().
			WithDescription("Unexpected return").
			WithStatus(common.StatusInvalidResponse)
	}
	if len(users) == 0 {
		return nil, common.NewResponse().
			WithDescription("Not found").
			WithStatus(common.StatusNotFound)
	}

	return users[0], nil
}
func (ops *Operations) SelectUserById(id int) (user *models.Users, res *common.Response) {
	var users []*models.Users

	res = ops.Query(
		models.UsersTable.Requests.SelectOneById.
			WithArgs(id).
			WithRowsScanner(scan.Users(&users)),
	)
	if res != nil {
		return user, res
	}
	if len(users) > 1 {
		return user, common.NewResponse().
			WithDescription("Unexpected return").
			WithStatus(common.StatusInvalidResponse)
	}
	if len(users) == 0 {
		return user, common.NewResponse().
			WithDescription("Not found").
			WithStatus(common.StatusNotFound)
	}

	return users[0], nil
}
func (ops *Operations) SelectUsersByInstitutionId(id int) ([]*models.Users, *transactions.Transaction) {
	tx := transactions.NewTransaction(nil)
	tx.Response = common.NewResponse().
		WithStatus(common.StatusUnregisteredMethod).
		WithDescription("Not implemented")
	return nil, tx
}
func (ops *Operations) SelectInstitutionByName(name string) (inst *models.Institutions, res *common.Response) {
	var insts []*models.Institutions

	res = ops.Query(
		models.InstitutionsTable.Requests.SelectOneByName.
			WithArgs(name).
			WithRowsScanner(scan.Institutions(&insts)),
	)
	if res != nil {
		return inst, res
	}

	if len(insts) > 1 {
		return inst, common.NewResponse().
			WithDescription("Unexpected return").
			WithStatus(common.StatusInvalidResponse)
	}
	if len(insts) == 0 {
		return inst, common.NewResponse().
			WithDescription("Not found").
			WithStatus(common.StatusNotFound)
	}

	return insts[0], nil
}
func (ops *Operations) SelectInstitutionById(id int) (inst *models.Institutions, res *common.Response) {
	var insts []*models.Institutions = []*models.Institutions{}
	res = ops.Query(
		models.InstitutionsTable.Requests.SelectOneById.
			WithArgs(id).
			WithRowsScanner(scan.Institutions(&insts)),
	)

	if res != nil {
		return nil, res
	}
	if len(insts) != 1 {
		return nil, common.NewResponse().
			WithDescription("Unexpected return").
			WithStatus(common.StatusInvalidResponse)
	}
	return insts[0], nil
}
func (ops *Operations) InsertManyUsers(requests ...*models.CreateUsers) (users []*models.Users, tx *transactions.Transaction) {
	tx = ops.Start()
	if tx.Response != nil {
		return users, tx
	}

	var args []any = []any{}
	for _, request := range requests {
		args = append(args, time.Now(), time.Now(), request.Name, request.Email, request.Password, request.InstitutionId, request.Role)
	}

	tx = tx.Query(models.UsersTable.Requests.InsertMany.
		WithArgs(args...).
		WithRowsScanner(scan.Users(&users)),
	)
	if tx.Response != nil {
		return users, tx
	}

	if len(users) != len(requests) {
		tx = transactions.NewTransaction(nil)
		tx.Response = common.NewResponse().
			WithDescription("Users created incorrectly").
			WithStatus(common.StatusNotEqual)
		return users, tx
	}

	return users, tx
}
func (ops *Operations) InsertInstitutions(requests ...*InstitutionRequest) (users []*models.Users, tx *transactions.Transaction) {
	tx = ops.Start()
	if tx.Response != nil {
		return users, tx
	}

	var args []any
	for _, request := range requests {
		args = append(args, time.Now(), time.Now(), request.Name)
	}

	var ids []int
	tx = tx.Query(models.InstitutionsTable.Requests.InsertMany.
		WithArgs(args...).
		WithRowsScanner(scan.InstitutionsIds(&ids)),
	)
	if tx.Response != nil {
		return users, tx
	}

	args = []any{}
	for _, request := range requests {
		args = append(args, time.Now(), time.Now(), "Administrator", request.Email, request.Password, ids[0], models.Admin)
	}
	tx = tx.Query(models.UsersTable.Requests.InsertMany.
		WithArgs(args...).
		WithRowsScanner(scan.Users(&users)),
	)
	if tx.Response != nil {
		return users, tx
	}

	return users, tx
}
func (ops *Operations) InsertNotifications(transaction *transactions.Transaction, requests ...*NotificationRequest) (tx *transactions.Transaction) {
	tx = ops.Start()
	if tx.Response != nil {
		return tx
	}

	var args []any = []any{}
	for _, request := range requests {
		args = append(args, time.Now(), time.Now(), request.Title, request.Description, request.IssuerId)
	}

	tx = tx.Query(
		models.NotificationsTable.Requests.InsertMany.
			WithArgs(args...),
	)
	if tx.Response != nil {
		return tx
	}

	args = []any{}
	for _, request := range requests {
		args = append(args, request.TargetId, request.TargetRole, request.IssuerId)
	}

	tx = tx.Query(
		models.UsersNotificationsTable.Requests.InsertMany.
			WithArgs(args...),
	)
	if tx.Response != nil {
		return tx
	}

	// Not done
	if true {
		panic("InsertNotifications implemented partially")
	}
	return tx
}
func (ops *Operations) DeleteUserById(id int) (tx *transactions.Transaction) {
	tx = ops.Start()
	if tx.Response != nil {
		return tx
	}

	tx.Response = tx.Query(models.UsersTable.Requests.DeleteOneById.WithArgs(id)).Response
	if tx.Response != nil {
		return tx
	}

	return tx
}
func (ops *Operations) DeleteUserByEmail(email string) (tx *transactions.Transaction) {
	tx = ops.Start()
	if tx.Response != nil {
		return tx
	}

	tx.Response = tx.Query(
		models.UsersTable.Requests.DeleteOneByEmail.
			WithArgs(email),
	).Response
	if tx.Response != nil {
		return tx
	}

	return tx
}
func (ops *Operations) DeleteInstitutionByName(transaction *transactions.Transaction, name string) (tx *transactions.Transaction) {
	tx = ops.Start()
	if tx.Response != nil {
		return tx
	}

	tx.Response = tx.Query(
		models.InstitutionsTable.Requests.DeleteOneByName.
			WithArgs(name),
	).Response
	if tx.Response != nil {
		return tx
	}

	return tx
}
func (ops *Operations) DeleteInstitutionById(transaction *transactions.Transaction, id int) (tx *transactions.Transaction) {
	tx = ops.Start()
	if tx.Response != nil {
		return tx
	}

	tx.Response = tx.Query(
		models.InstitutionsTable.Requests.DeleteOneById.
			WithArgs(id),
	).Response
	if tx.Response != nil {
		return tx
	}

	return tx
}
