package scan

import (
	"database/sql"

	"github.com/Noeeekr/singullar/server/common"
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/transactions"
)

/*
// If the return value is nil scan only check for errors on scan

	func UsersIds(ids *[]int) transactions.RequestRowsScanner {

		return func(Rows *sql.Rows, throwOnNotFound, throwOnFound) *common.Response {

			if ids == nil {
				return common.NewResponse().
					WithDescription("Cannot scan to nil pointer").
					WithStatus(common.StatusInvalidSyntax)
			}

			for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("Database found.").
					WithStatus(common.StatusFound)
			}

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
					if !throwErrorOnFound {
						return common.NewResponse().
							WithDescription("Database not found.").
							WithStatus(common.StatusNotFound)
					}

			return nil
		}
	}
*/

func DatabaseTypes(names *[]string) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if names == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("Database type already exists").
					WithStatus(common.StatusFound)
			}
			if err := Rows.Scan(&name); err != nil {
				return common.NewResponse().
					WithDescription(err.Error()).
					WithStatus(common.StatusInternalError)
			}
			*names = append(*names, name)
		}

		if Rows.Err() != nil {
			return common.NewResponse().
				WithDescription(Rows.Err().Error()).
				WithStatus(common.StatusInternalError)
		}
		if !throwErrorOnFound {
			if len(*names) == 0 {
				return common.NewResponse().
					WithDescription("Database type not found.").
					WithStatus(common.StatusNotFound)
			}
		}

		return nil
	}
}

func DatabaseTables(names *[]string) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if names == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("Database table already exists").
					WithStatus(common.StatusFound)
			}
			if err := Rows.Scan(&name); err != nil {
				return common.NewResponse().
					WithDescription(err.Error()).
					WithStatus(common.StatusInternalError)
			}
			*names = append(*names, name)
		}

		if Rows.Err() != nil {
			return common.NewResponse().
				WithDescription(Rows.Err().Error()).
				WithStatus(common.StatusInternalError)
		}
		if !throwErrorOnFound {
			if len(*names) == 0 {
				return common.NewResponse().
					WithDescription("Database table not found.").
					WithStatus(common.StatusNotFound)
			}
		}

		return nil
	}
}

func DatabaseUsers(names *[]string) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if names == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("Database user already exists").
					WithStatus(common.StatusFound)
			}
			if err := Rows.Scan(&name); err != nil {
				return common.NewResponse().
					WithDescription(err.Error()).
					WithStatus(common.StatusInternalError)
			}
			*names = append(*names, name)
		}

		if Rows.Err() != nil {
			return common.NewResponse().
				WithDescription(Rows.Err().Error()).
				WithStatus(common.StatusInternalError)
		}
		if !throwErrorOnFound {
			if len(*names) == 0 {
				return common.NewResponse().
					WithDescription("Database user not found.").
					WithStatus(common.StatusNotFound)
			}
		}

		return nil
	}
}
func DatabaseNames(names *[]string) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if names == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("Database found.").
					WithStatus(common.StatusFound)
			}
			if Rows.Scan(&name) != nil {
				return common.NewResponse().
					WithDescription(Rows.Err().Error()).
					WithStatus(common.StatusFailedTransaction)
			}
			*names = append(*names, name)
		}
		if Rows.Err() != nil {
			return common.NewResponse().WithDescription(Rows.Err().Error()).WithStatus(common.StatusFailedTransaction)
		}
		if !throwErrorOnFound {
			if len(*names) == 0 {
				return common.NewResponse().
					WithDescription("Database not found.").
					WithStatus(common.StatusNotFound)
			}
		}
		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func Notifications(notifications *[]*models.Notifications) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if notifications == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("Notification found.").
					WithStatus(common.StatusFound)
			}
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
		if !throwErrorOnFound {
			if len(*notifications) == 0 {
				return common.NewResponse().
					WithDescription("Notification not found.").
					WithStatus(common.StatusNotFound)
			}
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func UsersIds(ids *[]int) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if ids == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("User found.").
					WithStatus(common.StatusFound)
			}
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
		if !throwErrorOnFound {
			if len(*ids) == 0 {
				return common.NewResponse().
					WithDescription("User not found.").
					WithStatus(common.StatusNotFound)
			}
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func UsersEmail(emails *[]string) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if emails == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("User found.").
					WithStatus(common.StatusFound)
			}
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
		if !throwErrorOnFound {
			if len(*emails) == 0 {
				return common.NewResponse().
					WithDescription("User not found.").
					WithStatus(common.StatusNotFound)
			}
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func Users(users *[]*models.Users) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if users == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("User found.").
					WithStatus(common.StatusFound)
			}
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
		if !throwErrorOnFound {
			if len(*users) == 0 {
				return common.NewResponse().
					WithDescription("User not found.").
					WithStatus(common.StatusNotFound)
			}
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func InstitutionsIds(ids *[]int) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if ids == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("Institution found.").
					WithStatus(common.StatusFound)
			}
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
		if !throwErrorOnFound {
			if len(*ids) == 0 {
				return common.NewResponse().
					WithDescription("Institution not found.").
					WithStatus(common.StatusNotFound)
			}
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func Institutions(institutions *[]*models.Institutions) transactions.RequestRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) *common.Response {
		if institutions == nil {
			return common.NewResponse().
				WithDescription("Cannot scan to nil pointer").
				WithStatus(common.StatusInvalidSyntax)
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return common.NewResponse().
					WithDescription("Institution found.").
					WithStatus(common.StatusFound)
			}
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
		if !throwErrorOnFound {
			if len(*institutions) == 0 {
				return common.NewResponse().
					WithDescription("Institution not found.").
					WithStatus(common.StatusNotFound)
			}
		}

		return nil
	}
}
