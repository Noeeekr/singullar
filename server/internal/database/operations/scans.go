package operations

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

func scanUsersIds(ids *[]int) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *common.Response {
		if ids == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			var id int
			if Rows.Scan(&id) != nil {
				return common.NewResponse().
					WithDescription(Rows.Close().Error()).
					WithStatus(common.StatusFailedTransaction)
			}
			*ids = append(*ids, id)
		}

		if Rows.Err() != nil {
			return common.NewResponse().
				WithDescription(Rows.Err().Error()).
				WithStatus(common.StatusFailedTransaction)
		}

		return nil
	}
}
func scanUsersEmail(emails *[]string) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *common.Response {
		if emails == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			var email string
			if Rows.Scan(&email) != nil {
				return common.NewResponse().
					WithDescription(Rows.Err().Error()).
					WithStatus(common.StatusFailedTransaction)
			}
			*emails = append(*emails, email)
		}

		if Rows.Err() != nil {
			return common.NewResponse().
				WithDescription(Rows.Err().Error()).
				WithStatus(common.StatusFailedTransaction)
		}

		return nil
	}
}
func scanUsers(users *[]*models.Users) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *common.Response {
		if users == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			u := models.Users{}
			err := Rows.Scan(&u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Name, &u.Email, &u.Password, &u.InstitutionId, &u.Role, &u.Id, &u.ProfilePicture)
			if err != nil {
				return common.NewResponse().
					WithDescription(Rows.Err().Error()).
					WithStatus(common.StatusFailedTransaction)
			}
			*users = append(*users, &u)
		}

		if Rows.Err() != nil {
			return common.NewResponse().
				WithDescription(Rows.Err().Error()).
				WithStatus(common.StatusFailedTransaction)
		}

		return nil
	}
}

func scanInstitutionsIds(ids *[]int) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *common.Response {
		if ids == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			var id int
			if Rows.Scan(&id) != nil {
				return common.NewResponse().
					WithDescription(Rows.Err().Error()).
					WithStatus(common.StatusFailedTransaction)
			}
			*ids = append(*ids, id)
		}

		if Rows.Err() != nil {
			return common.NewResponse().
				WithDescription(Rows.Err().Error()).
				WithStatus(common.StatusFailedTransaction)
		}

		return nil
	}
}

func scanInstitutions(institutions *[]*models.Institutions) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *common.Response {
		if institutions == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			i := models.Institutions{}
			err := Rows.Scan(&i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Name, &i.Id)
			if err != nil {
				return common.NewResponse().
					WithDescription(Rows.Err().Error()).
					WithStatus(common.StatusFailedTransaction)
			}
			*institutions = append(*institutions, &i)
		}

		if Rows.Err() != nil {
			return common.NewResponse().
				WithDescription(Rows.Err().Error()).
				WithStatus(common.StatusFailedTransaction)
		}

		return nil
	}
}
