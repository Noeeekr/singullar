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
				return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
			}

			for Rows.Next() {

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
						return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
					}
					if !throwErrorOnFound {
						return borm.ErrorDescription(borm.ErrNotFound, "Database not found.")
					}

			return nil
		}
	}
*/

func DatabaseTypes(names *[]string) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if names == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if err := Rows.Scan(&name); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if len(*names) == 0 {
			return false, nil
		}

		return true, nil
	}
}

func DatabaseTables(names *[]string) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if names == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if err := Rows.Scan(&name); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*names) == 0 {
			return false, nil
		}

		return true, nil
	}
}

func DatabaseUsers(names *[]string) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if names == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if err := Rows.Scan(&name); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if len(*names) == 0 {
			return false, nil
		}
		return true, nil
	}
}
func DatabaseNames(names *[]string) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if names == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}
		defer Rows.Close()

		var name string
		for Rows.Next() {
			if err := Rows.Scan(&name); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}
		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*names) == 0 {
			return false, nil
		}
		return true, nil
	}
}

// If the return value is nil scan only check for errors on scan
func NotificationContents(notifications *[]*models.NotificationContents) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if notifications == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			var n models.NotificationContents
			if err := Rows.Scan(&n.CreatedAt, &n.UpdatedAt, &n.DeletedAt, &n.Id, &n.IssuerId, &n.Title, &n.Description); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*notifications = append(*notifications, &n)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*notifications) == 0 {
			return false, nil
		}

		return true, nil
	}
}

func Notifications(detailedNotifications *[]*models.Notifications) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if detailedNotifications == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			var n models.Notifications
			if err := Rows.Scan(&n.CreatedAt, &n.UpdatedAt, &n.DeletedAt, &n.TargetId, &n.TargetName, &n.IssuerId, &n.Title, &n.Description); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*detailedNotifications = append(*detailedNotifications, &n)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*detailedNotifications) == 0 {
			return false, nil
		}
		return true, nil
	}
}

func Integers(ids *[]int) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if ids == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			var id int
			if err := Rows.Scan(&id); err != nil {
				return false, borm.ErrorDescription(borm.ErrFailedTransaction, err.Error())
			}
			*ids = append(*ids, id)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*ids) == 0 {
			return false, nil
		}
		return true, nil
	}
}

// If the return value is nil scan only check for errors on scan
func UsersEmail(emails *[]string) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if emails == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			var email string
			if err := Rows.Scan(&email); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*emails = append(*emails, email)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if len(*emails) == 0 {
			return false, nil
		}
		return true, nil
	}
}

func Users(users *[]*models.Users) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if users == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			u := models.Users{}
			err := Rows.Scan(&u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Name, &u.Email, &u.Password, &u.InstitutionId, &u.Role, &u.Id, &u.ProfilePicture, &u.Segment)
			if err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*users = append(*users, &u)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*users) == 0 {
			return false, nil
		}

		return true, nil
	}
}

// If the return value is nil scan only check for errors on scan
func InstitutionsIds(ids *[]int) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if ids == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			var id int
			if err := Rows.Scan(&id); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*ids = append(*ids, id)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*ids) == 0 {
			return false, nil
		}
		return true, nil
	}
}

// If the return value is nil scan only check for errors on scan
func Institutions(institutions *[]*models.Institutions) borm.ReturnScanner {
	return func(Rows *sql.Rows) (bool, error) {
		if institutions == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for Rows.Next() {
			i := models.Institutions{}
			err := Rows.Scan(&i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Name, &i.Id)
			if err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())

			}
			*institutions = append(*institutions, &i)
		}

		if err := Rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		return len(*institutions) != 0, nil
	}
}
func Classes(classes *[]*models.Classes) borm.ReturnScanner {
	return func(rows *sql.Rows) (bool, error) {
		if classes == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for rows.Next() {
			class := models.Classes{}
			err := rows.Scan(&class.CreatedAt, &class.UpdatedAt, &class.DeletedAt, &class.Name, &class.Segment, &class.Series, &class.InstitutionId, &class.Id, &class.TeacherId)
			if err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*classes = append(*classes, &class)
		}

		if err := rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		return len(*classes) != 0, nil
	}
}
