package query

import (
	"database/sql"
	"errors"

	"github.com/Noeeekr/singullar/server/pkg/pim/models"
	"github.com/Noeeekr/singullar/server/pkg/pim/transaction"
)

func ScanUsers(users *[]*models.Users) transaction.QueryReturnFunc {
	return func(Rows *sql.Rows) error {
		return nil
	}
}

func ScanInstitutionsIds(ids *[]int) transaction.QueryReturnFunc {
	return func(Rows *sql.Rows) error {
		if ids == nil {
			return errors.New(" Unable to scan to nil. ")
		}

		if Rows.Err() != nil {
			return Rows.Err()
		}

		var id int
		for Rows.Next() {
			err := Rows.Scan(&id)
			if err != nil {
				return Rows.Err()
			}
		}

		if Rows.Err() != nil {
			return Rows.Err()
		}

		*ids = append(*ids, id)

		return nil
	}
}

func ScanInstitutions(institutions *[]*models.Institutions) transaction.QueryReturnFunc {
	return func(Rows *sql.Rows) error {
		if institutions == nil {
			return errors.New(" Unable to scan to nil. ")
		}

		if Rows.Err() != nil {
			return Rows.Err()
		}

		i := models.Institutions{}
		for Rows.Next() {
			err := Rows.Scan(&i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Name, &i.Id)
			if err != nil {
				return Rows.Err()
			}
		}

		if Rows.Err() != nil {
			return Rows.Err()
		}

		*institutions = append(*institutions, &i)

		return nil
	}
}
