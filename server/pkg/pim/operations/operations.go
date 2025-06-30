package operations

import (
	"database/sql"
	"time"

	"github.com/Noeeekr/singullar/server/pkg/pim/models"
	"github.com/Noeeekr/singullar/server/pkg/pim/transactions"
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

func (ops *Operations) InsertUser(name, email, password string, institution_id int, role models.UserRole) (email_ string, res *transactions.Response) {
	tx, res := ops.tx.Start()
	if res != nil {
		return email, res
	}

	var emails []string

	res = tx.Query(
		models.TablesInfo.Users.Requests.InsertOne.
			WithArgs(time.Now(), time.Now(), name, email, password, institution_id, role).
			WithScanFunc(scanUsersEmail(&emails)),
	)
	if res != nil {
		return email, res
	}
	if len(emails) != 1 {
		return email, res.SetDescription("Unexpected return").SetStatus(transactions.StatusInvalidResponse)
	}

	return email, tx.Commit()
}

func (ops *Operations) InsertInstitution(name string) (id int, res *transactions.Response) {
	tx, res := ops.tx.Start()
	if res != nil {
		return id, res
	}

	var ids []int

	res = tx.Query(models.TablesInfo.Institutions.Requests.InsertOne.
		WithArgs(time.Now(), time.Now(), name).
		WithScanFunc(scanInstitutionsIds(&ids)),
	)
	if res != nil {
		return id, res
	}
	if len(ids) != 1 {
		return 0, res.SetDescription("Unexpected return").SetStatus(transactions.StatusInvalidResponse)
	}

	return ids[0], tx.Commit()
}
func (ops *Operations) SelectInstitutionByName(name string) (inst *models.Institutions, res *transactions.Response) {
	var insts []*models.Institutions

	res = ops.tx.Query(
		models.TablesInfo.Institutions.Requests.SelectOneByName.
			WithArgs(name).
			WithScanFunc(scanInstitutions(&insts)),
	)
	if res != nil {
		return inst, res
	}

	if len(insts) != 1 || insts[0] == nil {
		return inst, res.SetDescription("Unexpected return").SetStatus(transactions.StatusInvalidResponse)
	}

	return insts[0], nil
}
func (ops *Operations) SelectInstitutionById(id int) (inst *models.Institutions, res *transactions.Response) {
	var insts []*models.Institutions

	res = ops.tx.Query(
		models.TablesInfo.Institutions.Requests.SelectOneById.
			WithArgs(id).
			WithScanFunc(scanInstitutions(&insts)),
	)
	if res != nil {
		return nil, res
	}

	if len(insts) != 1 || insts[0] == nil {
		return nil, res.SetDescription("Unexpected return").SetStatus(transactions.StatusInvalidResponse)
	}
	return insts[0], nil
}
func (ops *Operations) DeleteInstitutionByName(name string) (res *transactions.Response) {
	tx, res := ops.tx.Start()
	if res != nil {
		return res
	}

	res = tx.Query(
		models.TablesInfo.Institutions.Requests.DeleteOneByName.
			WithArgs(name),
	)
	if res != nil {
		return res
	}

	return tx.Commit()
}
func (ops *Operations) DeleteInstitutionById(id int) (res *transactions.Response) {
	tx, res := ops.tx.Start()
	if res != nil {
		return res
	}

	res = tx.Query(
		models.TablesInfo.Institutions.Requests.DeleteOneById.
			WithArgs(id),
	)
	if res != nil {
		return res
	}

	return tx.Commit()
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
