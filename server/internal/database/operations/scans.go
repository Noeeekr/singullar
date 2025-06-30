package operations

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

func scanUsersEmail(emails *[]string) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *transactions.Response {
		if emails == nil {
			return transactions.NewResponse().
				SetDescription("Cannot scan to nil pointer").
				SetStatus(transactions.StatusInvalidSyntax)
		}

		var email string
		for Rows.Next() {
			if Rows.Scan(&email) != nil {
				return transactions.NewResponse().
					SetDescription(Rows.Err().Error()).
					SetStatus(transactions.StatusFailedTransaction)
			}
		}

		if Rows.Err() != nil {
			return transactions.NewResponse().
				SetDescription(Rows.Err().Error()).
				SetStatus(transactions.StatusFailedTransaction)
		}

		*emails = append(*emails, email)

		return nil
	}
}
func scanUsers(users *[]*models.Users) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *transactions.Response {
		if users == nil {
			return transactions.NewResponse().
				SetDescription("Cannot scan to nil pointer").
				SetStatus(transactions.StatusInvalidSyntax)
		}

		u := models.Users{}
		for Rows.Next() {
			err := Rows.Scan(&u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Name, &u.Email, &u.Password, &u.InstitutionId, &u.Role, &u.Id, &u.ProfilePicture)
			if err != nil {
				return transactions.NewResponse().
					SetDescription(Rows.Err().Error()).
					SetStatus(transactions.StatusFailedTransaction)
			}
		}

		if Rows.Err() != nil {
			return transactions.NewResponse().
				SetDescription(Rows.Err().Error()).
				SetStatus(transactions.StatusFailedTransaction)
		}

		*users = append(*users, &u)

		return nil
	}
}

func scanInstitutionsIds(ids *[]int) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *transactions.Response {
		if ids == nil {
			return transactions.NewResponse().
				SetDescription("Cannot scan to nil pointer").
				SetStatus(transactions.StatusInvalidSyntax)
		}

		var id int
		for Rows.Next() {
			if Rows.Scan(&id) != nil {
				return transactions.NewResponse().
					SetDescription(Rows.Err().Error()).
					SetStatus(transactions.StatusFailedTransaction)
			}
		}

		if Rows.Err() != nil {
			return transactions.NewResponse().
				SetDescription(Rows.Err().Error()).
				SetStatus(transactions.StatusFailedTransaction)
		}

		*ids = append(*ids, id)

		return nil
	}
}

func scanInstitutions(institutions *[]*models.Institutions) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *transactions.Response {
		if institutions == nil {
			return transactions.NewResponse().
				SetDescription("Cannot scan to nil pointer").
				SetStatus(transactions.StatusInvalidSyntax)
		}

		i := models.Institutions{}
		for Rows.Next() {
			err := Rows.Scan(&i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Name, &i.Id)
			if err != nil {
				return transactions.NewResponse().
					SetDescription(Rows.Err().Error()).
					SetStatus(transactions.StatusFailedTransaction)
			}
		}

		if Rows.Err() != nil {
			return transactions.NewResponse().
				SetDescription(Rows.Err().Error()).
				SetStatus(transactions.StatusFailedTransaction)
		}

		*institutions = append(*institutions, &i)

		return nil
	}
}
