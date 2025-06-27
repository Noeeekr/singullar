package managers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Noeeekr/singullar/server/pkg/pim/models"
)

type QueryManager struct {
	*TransactionManager
}

func NewQueryManager(db *sql.DB) *QueryManager {
	return &QueryManager{
		&TransactionManager{
			db: db,
		},
	}
}

func (m *QueryManager) SelectInstitution(id int) (*models.Institutions, *Response) {
	res := &Response{}
	i := &models.Institutions{}

	query := &Query{
		Query:   models.InstitutionsTable.Queries.SelectOne,
		Returns: true,
		Args:    []any{id},
	}

	tx, err := m.StartTransation()
	if err != nil {
		res.Description = err.Error()
		return i, res
	}

	result := tx.Query(query)
	if result.Status != StatusSuccess {
		return i, result.Response
	}

	for result.Rows.Next() {
		err := result.Rows.Scan(&i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Name, &i.Id)
		if err != nil {
			res.Status = StatusFailedTransaction
			res.Description = err.Error()
			return i, res
		}
	}

	return i, tx.Commit().Response
}

func (m *QueryManager) InsertInstitution(name string) (id int, res *Response) {
	res = &Response{}

	query := &Query{
		Query:   models.InstitutionsTable.Queries.InsertOne,
		Returns: true,
		Args:    []any{time.Now(), time.Now(), name},
	}

	tx, err := m.StartTransation()
	if err != nil {
		res.Description = err.Error()
		return id, res
	}

	result := tx.Query(query)
	if result.Status != StatusSuccess {
		return id, result.Response
	}

	for result.Rows.Next() {
		err := result.Rows.Scan(&id)
		if err != nil {
			res.Status = StatusFailedTransaction
			res.Description = err.Error()
			return id, res
		}
	}

	result = tx.Commit()

	return id, result.Response
}

func (m *QueryManager) DeleteInstitution(id int) (res *Response) {
	res = &Response{}

	query := &Query{
		Query:   fmt.Sprintf("DELETE FROM %s WHERE id = %d", models.InstitutionsTable.Name(), id),
		Returns: false,
	}

	tx, err := m.StartTransation()
	if err != nil {
		res.Description = err.Error()
		return res
	}

	result := tx.Query(query)
	if result.Status != StatusSuccess {
		return result.Response
	}

	return tx.Commit().Response
}

func (m *QueryManager) InsertUser(name, email, password string, institution_id int, role models.UserRole) (Email string, res *Response) {
	res = &Response{}

	tx, err := m.StartTransation()
	if err != nil {
		res.Description = err.Error()
		return Email, res
	}

	// Check if user exists
	query := &Query{
		Query:   models.UsersTable.Queries.SelectOne,
		Returns: true,
		Args:    []any{email},
	}

	result := tx.Query(query)
	if result.Status != StatusSuccess {
		return Email, result.Response
	}
	if result.Rows.Next() {
		return Email, &Response{
			Status:      StatusAlreadyExists,
			Description: "User already exists",
		}
	}

	// Insert
	query = &Query{
		Query:   models.UsersTable.Queries.InsertOne,
		Returns: true,
		Args:    []any{time.Now(), time.Now(), name, email, password, institution_id, role},
	}

	result = tx.Query(query)
	if result.Status != StatusSuccess {
		return Email, result.Response
	}

	for result.Rows.Next() {
		err := result.Rows.Scan(&Email)
		if err != nil {
			res.Status = StatusFailedTransaction
			res.Description = err.Error()
			return Email, res
		}
	}

	result = tx.Commit()

	return Email, result.Response
}

func (m *QueryManager) SelectUser(email string) (*models.Users, *Response) {
	res := &Response{}
	u := &models.Users{}

	query := &Query{
		Query:   models.UsersTable.Queries.SelectOne,
		Returns: true,
		Args:    []any{email},
	}

	tx, err := m.StartTransation()
	if err != nil {
		res.Description = err.Error()
		return u, res
	}

	result := tx.Query(query)
	if result.Status != StatusSuccess {
		return u, result.Response
	}

	for result.Rows.Next() {
		err := result.Rows.Scan(&u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Name, &u.Email, &u.Password, &u.InstitutionId, &u.Role, &u.Id, &u.ProfilePicture)
		if err != nil {
			res.Status = StatusFailedTransaction
			res.Description = err.Error()
			return u, res
		}
	}

	return u, tx.Commit().Response
}

func (m *QueryManager) DeleteUser(email string) *Response {
	res := &Response{}

	query := &Query{
		Query:   fmt.Sprintf("DELETE FROM %s WHERE email = '%s'", models.UsersTable.Name(), email),
		Returns: false,
	}

	tx, err := m.StartTransation()
	if err != nil {
		res.Description = err.Error()
		return res
	}

	result := tx.Query(query)
	if result.Status != StatusSuccess {
		return result.Response
	}

	return tx.Commit().Response
}

/*
	user: Delete Institution
		sends: password && email
			happens:
				Select User where Email
				If User.PasswordHash == Password
					Delete User.Institution
*/
