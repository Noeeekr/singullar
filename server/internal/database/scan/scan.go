package scan

import (
	"database/sql"

	"github.com/Noeeekr/borm"
	"github.com/Noeeekr/singullar/server/internal/database/models"
)

/*
// If the return value is nil scan only check for errors on scan

	func UsersIds(ids *[]int) borm.QueryrowsScanner {

		return func(rows *sql.Rows, throwOnNotFound, throwOnFound) error {

			if ids == nil {
				return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
			}

			for rows.Next() {

				if throwOnFound {
					return borm.ErrorDescription(StatusFound), "Found"				}

				var id int
				if err := rows.Scan(&id); err != nil {
					return borm.ErrorDescription(borm.ErrFailedTransaction, ).Error())
				}
				*ids = append(*ids, id)
			}

			if len(ids) == 0 {
				if throwAtNotFound {
					return borm.ErrorDescription(StatusNotFound), "NotFound"				}
			}

			if err := rows.Err(); err != nil {
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
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if names == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		var name string
		for rows.Next() {
			if err := rows.Scan(&name); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}

		if err := rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if len(*names) == 0 {
			return false, nil
		}

		return true, nil
	}
}

func DatabaseTables(names *[]string) borm.ReturnScanner {
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if names == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		var name string
		for rows.Next() {
			if err := rows.Scan(&name); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}

		if err := rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*names) == 0 {
			return false, nil
		}

		return true, nil
	}
}

func DatabaseUsers(names *[]string) borm.ReturnScanner {
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if names == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		var name string
		for rows.Next() {
			if err := rows.Scan(&name); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}

		if err := rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if len(*names) == 0 {
			return false, nil
		}
		return true, nil
	}
}
func DatabaseNames(names *[]string) borm.ReturnScanner {
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if names == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		var name string
		for rows.Next() {
			if err := rows.Scan(&name); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*names = append(*names, name)
		}
		if err := rows.Err(); err != nil {
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
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if notifications == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for rows.Next() {
			var n models.NotificationContents
			if err := rows.Scan(&n.CreatedAt, &n.UpdatedAt, &n.DeletedAt, &n.Id, &n.IssuerId, &n.Title, &n.Description); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*notifications = append(*notifications, &n)
		}

		if err := rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*notifications) == 0 {
			return false, nil
		}

		return true, nil
	}
}

func Notifications(detailedNotifications *[]*models.Notifications) borm.ReturnScanner {
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if detailedNotifications == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for rows.Next() {
			var n models.Notifications
			if err := rows.Scan(&n.CreatedAt, &n.UpdatedAt, &n.DeletedAt, &n.TargetId, &n.TargetName, &n.IssuerId, &n.Title, &n.Description); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*detailedNotifications = append(*detailedNotifications, &n)
		}

		if err := rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		if len(*detailedNotifications) == 0 {
			return false, nil
		}
		return true, nil
	}
}

func Integers(ids *[]int) borm.ReturnScanner {
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if ids == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				return false, borm.ErrorDescription(borm.ErrFailedTransaction, err.Error())
			}
			*ids = append(*ids, id)
		}

		if err := rows.Err(); err != nil {
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
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if emails == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for rows.Next() {
			var email string
			if err := rows.Scan(&email); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*emails = append(*emails, email)
		}

		if err := rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		if len(*emails) == 0 {
			return false, nil
		}
		return true, nil
	}
}

func Users(users *[]*models.Users) borm.ReturnScanner {
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if users == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for rows.Next() {
			u := models.Users{}
			err := rows.Scan(&u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Name, &u.Email, &u.Password, &u.InstitutionId, &u.Role, &u.Id, &u.ProfilePicture, &u.Segment)
			if err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*users = append(*users, &u)
		}

		if err := rows.Err(); err != nil {
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
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if ids == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
			}
			*ids = append(*ids, id)
		}

		if err := rows.Err(); err != nil {
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
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
		if institutions == nil {
			return false, borm.ErrorDescription(borm.ErrSyntax, "Cannot scan to nil pointer")
		}

		for rows.Next() {
			i := models.Institutions{}
			err := rows.Scan(&i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Name, &i.Id)
			if err != nil {
				return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())

			}
			*institutions = append(*institutions, &i)
		}

		if err := rows.Err(); err != nil {
			return false, borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}

		return len(*institutions) != 0, nil
	}
}
func QuestionLists(lists *[]*models.QuestionLists) borm.ReturnScanner {
	return Scanner[models.QuestionLists](func(r *sql.Rows) error {
		list := models.QuestionLists{}
		err := r.Scan(&list.Id, &list.CreatedAt, &list.UpdatedAt, &list.InstitutionId, &list.Title, &list.SubjectId, &list.DifficultyLevel)
		if err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		*lists = append(*lists, &list)
		return nil
	}).On(lists)
}
func ExtendedQuestionLists(lists *[]*models.ExtendedQuestionList) borm.ReturnScanner {
	return Scanner[models.ExtendedQuestionList](func(r *sql.Rows) error {
		list := models.ExtendedQuestionList{}
		err := r.Scan(&list.Id, &list.CreatedAt, &list.UpdatedAt, &list.InstitutionId, &list.Title, &list.SubjectId, &list.SubjectName, &list.DifficultyLevel, &list.QuestionQuantity)
		if err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		*lists = append(*lists, &list)
		return nil
	}).On(lists)
}
func Questions(questions *[]*models.Questions) borm.ReturnScanner {
	return Scanner[models.Questions](func(r *sql.Rows) error {
		question := models.Questions{}
		err := r.Scan(&question.Id, &question.QuestionInstitutionId, &question.CreatedAt, &question.UpdatedAt,
			&question.QuestionTitle, &question.QuestionDescription, &question.QuestionShortDescription,
			&question.QuestionDifficultyLevel, &question.QuestionCorrectAlternative, &question.QuestionSubjectId)
		if err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		*questions = append(*questions, &question)
		return nil
	}).On(questions)
}
func Classes(classes *[]*models.Classes) borm.ReturnScanner {
	return func(rows *sql.Rows) (bool, error) {
		defer rows.Close()
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

func Subjects(subjects *[]*models.Subjects) borm.ReturnScanner {
	return Scanner[models.Subjects](func(row *sql.Rows) error {
		subject := models.Subjects{}
		err := row.Scan(&subject.Id, &subject.SubjectName, &subject.InstitutionId)
		if err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		*subjects = append(*subjects, &subject)
		return nil
	}).On(subjects)
}

func QuestionDifficulties(difficulties *[]*models.QuestionDifficulty) borm.ReturnScanner {
	return Scanner[models.QuestionDifficulty](func(r *sql.Rows) error {
		questionListDifficulty := models.QuestionDifficulty{}
		err := r.Scan(&questionListDifficulty.DifficultyName, &questionListDifficulty.DifficultyLevel)
		if err != nil {
			return borm.ErrorDescription(borm.ErrUnexpected, "Error while scanning rows", err.Error())
		}
		*difficulties = append(*difficulties, &questionListDifficulty)
		return nil
	}).On(difficulties)
}
