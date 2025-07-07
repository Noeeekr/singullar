package operations

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

type Operations struct {
	tx *transactions.TransactionManager
}

// Returns an instance of Operations. Operations contains methods that to make the most used transactions instantly.
func New(db *sql.DB) *Operations {
	return &Operations{
		tx: transactions.New(db),
	}
}

// returns the email
func (ops *Operations) InsertManyUsers(transaction *transactions.Transaction, requests ...*models.CreateUsers) (users []*models.Users, tx *transactions.Transaction) {
	if transaction == nil {
		tx = ops.tx.Start()
		if tx.Response != nil {
			return users, tx
		}
	} else {
		tx = transaction
	}

	var args []any = []any{}
	for _, request := range requests {
		args = append(args, time.Now(), time.Now(), request.Name, request.Email, request.Password, request.InstitutionId, request.Role)
	}

	tx = tx.Query(models.TablesInfo.Users.Requests.InsertMany.
		WithArgs(args...).
		WithScanFunc(scanUsers(&users)),
	)
	if tx.Response != nil {
		return users, tx
	}

	if len(users) != len(requests) {
		tx = transactions.NewTransaction()
		tx.Response = common.NewResponse().
			WithDescription("Users created incorrectly").
			WithStatus(common.StatusNotEqual)
		return users, tx
	}

	if transaction == nil {
		return users, tx.Commit()
	}
	return users, tx
}

func (ops *Operations) SelectUserByEmail(email string) (user *models.Users, res *common.Response) {
	var users []*models.Users

	res = ops.tx.Query(
		models.TablesInfo.Users.Requests.SelectOneByEmail.
			WithArgs(email).
			WithScanFunc(scanUsers(&users)),
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

	res = ops.tx.Query(
		models.TablesInfo.Users.Requests.SelectOneById.
			WithArgs(id).
			WithScanFunc(scanUsers(&users)),
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
	tx := transactions.NewTransaction()
	tx.Response = common.NewResponse().
		WithStatus(common.StatusUnregisteredMethod).
		WithDescription("Not implemented")
	return nil, tx
}
func (ops *Operations) DeleteUserById(transaction *transactions.Transaction, id int) (tx *transactions.Transaction) {
	if transaction == nil {
		tx = ops.tx.Start()
		if tx.Response != nil {
			return tx
		}
	} else {
		tx = transaction
	}

	tx.Response = tx.Query(models.TablesInfo.Users.Requests.DeleteOneById.WithArgs(id)).Response
	if tx.Response != nil {
		return tx
	}

	if transaction == nil {
		return tx.Commit()
	}
	return tx
}
func (ops *Operations) DeleteUserByEmail(transaction *transactions.Transaction, email string) (tx *transactions.Transaction) {
	if transaction == nil {
		tx = ops.tx.Start()
		if tx.Response != nil {
			return tx
		}
	} else {
		tx = transaction
	}

	tx.Response = tx.Query(
		models.TablesInfo.Users.Requests.DeleteOneByEmail.
			WithArgs(email),
	).Response
	if tx.Response != nil {
		return tx
	}

	if transaction == nil {
		return tx.Commit()
	}
	return tx
}

func (ops *Operations) InsertInstitution(transaction *transactions.Transaction, name, email, password string) (user *models.Users, tx *transactions.Transaction) {
	if transaction == nil {
		tx = ops.tx.Start()
		if tx.Response != nil {
			return user, tx
		}
	} else {
		tx = transaction
	}

	var ids []int

	tx.Response = tx.Query(models.TablesInfo.Institutions.Requests.InsertOne.
		WithArgs(time.Now(), time.Now(), name).
		WithScanFunc(scanInstitutionsIds(&ids)),
	).Response
	if tx.Response != nil {
		return
	}

	if len(ids) > 1 {
		tx.Response = common.NewResponse().
			WithDescription("Unexpected return").
			WithStatus(common.StatusInvalidResponse)
		return
	}
	if len(ids) == 0 {
		tx.Response = common.NewResponse().
			WithDescription("Empty response").
			WithStatus(common.StatusNotFound)
		return
	}

	fmt.Println(models.TablesInfo.Users.Requests.InsertMany.
		WithArgs(time.Now(), time.Now(), "Administrator", email, password, ids[0], models.Admin).Query,
	)
	var users []*models.Users
	tx.Response = tx.Query(models.TablesInfo.Users.Requests.InsertMany.
		WithArgs(time.Now(), time.Now(), "Administrator", email, password, ids[0], models.Admin).
		WithScanFunc(scanUsers(&users)),
	).Response
	if tx.Response != nil {
		return user, tx
	}

	if len(users) != 1 {
		tx.Response = common.NewResponse().
			WithDescription("Unexpected return").
			WithStatus(common.StatusInvalidResponse)
		return
	}
	if len(users) == 0 {
		tx.Response = common.NewResponse().
			WithDescription("Empty response").
			WithStatus(common.StatusNotFound)
		return
	}

	if transaction == nil {
		return users[0], tx.Commit()
	}
	return users[0], transaction
}

func (ops *Operations) SelectInstitutionByName(name string) (inst *models.Institutions, res *common.Response) {
	var insts []*models.Institutions

	res = ops.tx.Query(
		models.TablesInfo.Institutions.Requests.SelectOneByName.
			WithArgs(name).
			WithScanFunc(scanInstitutions(&insts)),
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

	res = ops.tx.Query(
		models.TablesInfo.Institutions.Requests.SelectOneById.
			WithArgs(id).
			WithScanFunc(scanInstitutions(&insts)),
	)

	if res != nil {
		return nil, res
	}
	if len(insts) != 1 {
		return nil, common.NewResponse().
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
func (ops *Operations) DeleteInstitutionByName(transaction *transactions.Transaction, name string) (tx *transactions.Transaction) {
	if transaction == nil {
		tx = ops.tx.Start()
		if tx.Response != nil {
			return tx
		}
	}

	tx.Response = tx.Query(
		models.TablesInfo.Institutions.Requests.DeleteOneByName.
			WithArgs(name),
	).Response
	if tx.Response != nil {
		return tx
	}

	if transaction == nil {
		return tx.Commit()
	}
	return tx
}
func (ops *Operations) DeleteInstitutionById(transaction *transactions.Transaction, id int) (tx *transactions.Transaction) {
	if transaction == nil {
		tx = ops.tx.Start()
		if tx.Response != nil {
			return tx
		}
	} else {
		tx = transaction
	}

	tx.Response = tx.Query(
		models.TablesInfo.Institutions.Requests.DeleteOneById.
			WithArgs(id),
	).Response
	if tx.Response != nil {
		return tx
	}

	if transaction == nil {
		return tx.Commit()
	}
	return tx
}

/*
utils.Insert(&Table{
		,
		Name: models.InstitutionsTable.Name(),
	})
	if len(institutions_ids) == 0 {
		test.Fatal("Failed Scan")
	}
	utils.Select(&Table{
		QueryInfo: models.InstitutionsTable.Queries.SelectOne.
			WithArgs(institutions_ids[0]).
			WithScanFunc(operations.ScanInstitutions(&institutions)),
		Name: models.InstitutionsTable.Name(),
	})
	if len(institutions) == 0 {
		test.Fatal("Failed Scan")
	}
	utils.Insert(&Table{
		Name: models.UsersTable.Name(),
		QueryInfo: models.UsersTable.Queries.InsertOne.
			WithArgs(time.Now(), time.Now(), "noeeekr", user_email, "123123123", institutions_ids[0], models.Admin).
			WithScanFunc(operations.ScanUsersEmails(&users_emails)),
	})
	if len(users_emails) == 0 {
		test.Fatal("Failed Scan")
	}
	utils.Select(&Table{
		Name: models.UsersTable.Name(),
		QueryInfo: models.UsersTable.Queries.SelectOne.
			WithArgs(users_emails[0]).
			WithScanFunc(operations.ScanUsers(&users)),
	})
	if len(users) == 0 {
		test.Fatal("Failed Scan")
	}
	utils.Delete(&Table{
		//QueryInfo: models.UsersTable.Queries.,
	})
Delete User <- Email
Delete Inst <- ID


*/
