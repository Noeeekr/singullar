package manager

import (
	"github.com/Noeeekr/singullar/server/internal/database/models"
	"github.com/Noeeekr/singullar/server/internal/database/scan"
)

func (ops *Operator) DeleteUserById(id int) error {
	query := models.TableUsers.Delete()
	query.Where(query.Field("id").IsEqual(id))
	err := ops.Do(query)
	if err != nil {
		return err
	}
	return err
}
func (ops *Operator) DeleteUserByEmail(email string) error {
	query := models.TableUsers.Delete()
	query.Where(query.Field("email").IsEqual(email))
	err := ops.Do(query)
	return err
}
func (ops *Operator) DeleteInstitutionByName(name string) error {
	var ids []int
	query := models.TableInstitutions.Delete()
	query.Where(query.Field("name").IsEqual(name)).
		Returning("id").
		Scanner(scan.Integers(&ids))
	err := ops.Do(query)

	return err
}
func (ops *Operator) DeleteInstitutionById(id int) error {
	query := models.TableInstitutions.Delete()
	query.Where(query.Field("id").IsEqual(id))
	err := ops.Do(query)
	return err
}

func (ops *Operator) DeleteNotificationsByIssuerId(ids ...int) error {
	for _, id := range ids {
		query := models.TableNotificationContents.Delete()
		query.Where(query.Field("id").IsEqual(id))
		err := ops.Do(query)
		if err != nil {
			return err
		}
	}
	return nil
}
