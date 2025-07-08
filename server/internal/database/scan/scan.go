package scan

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

/*
// If the return value is nil scan only check for errors on scan

	func UsersIds(ids *[]int) transactions.RequestReturnHandler {

		return func(Rows *sql.Rows, throwOnNotFound, throwOnFound) *common.Response {

			if ids == nil {
				return common.NewResponse().
					WithDescription("Cannot scan to nil pointer").
					WithStatus(common.StatusInvalidSyntax)
			}

			for Rows.Next() {

				if throwOnFound {
					return common.NewResponse()
						.WithDescription("Found").WithStatus(StatusFound)
				}

				var id int
				if Rows.Scan(&id) != nil {
					return common.NewResponse().
						WithDescription(Rows.Close().Error()).
						WithStatus(common.StatusFailedTransaction)
				}
				*ids = append(*ids, id)
			}

			if len(ids) == 0 {
				if throwAtNotFound {
					return common.NewResponse()
						.WithDescription("NotFound").WithStatus(StatusNotFound)
				}
			}

			if Rows.Err() != nil {
				return common.NewResponse().
					WithDescription(Rows.Err().Error()).
					WithStatus(common.StatusFailedTransaction)
			}

			return nil
		}
	}
*/
func DatabaseNames(names *[]string) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *common.Response {
		if names == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}
		defer Rows.Close()

		var name string
		if Rows.Next() {
			if Rows.Scan(&name) != nil {
				return common.NewResponse().
					WithDescription(Rows.Err().Error()).
					WithStatus(common.StatusFailedTransaction)
			}
			*names = append(*names, name)
			return common.NewResponse().WithDescription("Database already exists.").WithStatus(common.StatusFound)
		}
		if Rows.Err() != nil {
			return common.NewResponse().WithDescription(Rows.Err().Error()).WithStatus(common.StatusFailedTransaction)
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func Notifications(notifications *[]*models.Notifications) transactions.RequestReturnHandler {
	return func(Rows *sql.Rows) *common.Response {
		if notifications == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			var notification *models.Notifications
			if Rows.Scan(&notification) != nil {
				return common.NewResponse().
					WithDescription(Rows.Err().Error()).
					WithStatus(common.StatusFailedTransaction)
			}
			*notifications = append(*notifications, notification)
		}

		if Rows.Err() != nil {
			return common.NewResponse().
				WithDescription(Rows.Err().Error()).
				WithStatus(common.StatusFailedTransaction)
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func UsersIds(ids *[]int) transactions.RequestReturnHandler {
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

// If the return value is nil scan only check for errors on scan
func UsersEmail(emails *[]string) transactions.RequestReturnHandler {
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

// If the return value is nil scan only check for errors on scan
func Users(users *[]*models.Users) transactions.RequestReturnHandler {
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

// If the return value is nil scan only check for errors on scan
func InstitutionsIds(ids *[]int) transactions.RequestReturnHandler {
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

// If the return value is nil scan only check for errors on scan
func Institutions(institutions *[]*models.Institutions) transactions.RequestReturnHandler {
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
