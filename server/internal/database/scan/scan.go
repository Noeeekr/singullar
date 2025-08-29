package scan

import (
	"database/sql"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/database/models"
)

/*
// If the return value is nil scan only check for errors on scan

	func UsersIds(ids *[]int) borm.QueryRowsScanner {

		return func(Rows *sql.Rows, throwOnNotFound, throwOnFound) error {

			if ids == nil {
				return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
			}

			for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "Database found.")
			}

				if throwOnFound {
					return borm.ErrorDescription(StatusFound), "Found"				}

				var id int
				if err := Rows.Scan(&id); err != nil {
					return borm.ErrorDescription(borm.ErrFailedTransaction, ).Error())
				}
				*ids = append(*ids, id)
			}

			if len(ids) == 0 {
				if throwAtNotFound {
					return borm.ErrorDescription(StatusNotFound), "NotFound"				}
			}

			if err := Rows.Err(); err != nil {
						return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
					}
					if !throwErrorOnFound {
						return borm.ErrorDescription(borm.ErrNotFound, "Database not found.")
					}

			return nil
		}
	}
*/

func DatabaseTypes(names *[]string) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if names == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "Database type already exists")
			}
			if err := Rows.Scan(&name); err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*names) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "Database type not found.")
			}
		}

		return nil
	}
}

func DatabaseTables(names *[]string) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if names == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "Database table already exists")
			}
			if err := Rows.Scan(&name); err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*names) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "Database table not found.")
			}
		}

		return nil
	}
}

func DatabaseUsers(names *[]string) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if names == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "Database user already exists")
			}
			if err := Rows.Scan(&name); err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*names) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "Database user not found.")
			}
		}

		return nil
	}
}
func DatabaseNames(names *[]string) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if names == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "Database found.")
			}
			if err := Rows.Scan(&name); err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}
		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*names) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "Database not found.")
			}
		}
		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func NotificationContents(notifications *[]*models.NotificationContents) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if notifications == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "Notification found.")
			}
			var n models.NotificationContents
			if err := Rows.Scan(&n.CreatedAt, &n.UpdatedAt, &n.DeletedAt, &n.Id, &n.IssuerId, &n.Title, &n.Description); err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*notifications = append(*notifications, &n)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*notifications) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "Notification not found.")
			}
		}

		return nil
	}
}

func Notifications(detailedNotifications *[]*models.Notifications) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if detailedNotifications == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "Notification found.")
			}
			var n models.Notifications
			if err := Rows.Scan(&n.CreatedAt, &n.UpdatedAt, &n.DeletedAt, &n.TargetId, &n.TargetName, &n.IssuerId, &n.Title, &n.Description); err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*detailedNotifications = append(*detailedNotifications, &n)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*detailedNotifications) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "Notification not found.")
			}
		}

		return nil
	}
}

func Integers(ids *[]int) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if ids == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "User found.")
			}
			var id int
			if err := Rows.Scan(&id); err != nil {
				return borm.ErrorDescription(borm.ErrFailedTransaction, err.Error())
			}
			*ids = append(*ids, id)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*ids) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "User not found.")
			}
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func UsersEmail(emails *[]string) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if emails == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "User found.")
			}
			var email string
			if err := Rows.Scan(&email); err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*emails = append(*emails, email)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*emails) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "User not found.")
			}
		}

		return nil
	}
}

func Users(users *[]*models.Users) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if users == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "User found.")
			}
			u := models.Users{}
			err := Rows.Scan(&u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Name, &u.Email, &u.Password, &u.InstitutionId, &u.Role, &u.Id, &u.ProfilePicture, &u.Segment)
			if err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*users = append(*users, &u)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*users) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "User not found.")
			}
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func InstitutionsIds(ids *[]int) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if ids == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "Institution found.")
			}
			var id int
			if err := Rows.Scan(&id); err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*ids = append(*ids, id)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*ids) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "Institution not found.")
			}
		}

		return nil
	}
}

// If the return value is nil scan only check for errors on scan
func Institutions(institutions *[]*models.Institutions) borm.QueryRowsScanner {
	return func(Rows *sql.Rows, throwErrorOnFound bool) error {
		if institutions == nil {
			return borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			if throwErrorOnFound {
				return borm.ErrorDescription(borm.ErrFound, "Institution found.")
			}
			i := models.Institutions{}
			err := Rows.Scan(&i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Name, &i.Id)
			if err != nil {
				return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())

			}
			*institutions = append(*institutions, &i)
		}

		if err := Rows.Err(); err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if !throwErrorOnFound {
			if len(*institutions) == 0 {
				return borm.ErrorDescription(borm.ErrNotFound, "Institution not found.")
			}
		}

		return nil
	}
}
