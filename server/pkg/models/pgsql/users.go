package pgsql

import (
	"errors"

	"gorm.io/gorm"

	"github.com/noeeekr/sch-server/pkg/models"
)

// database model methods initializer
type UserModel struct {
	DB *gorm.DB
}

func (model *UserModel) Insert(user *models.Users) (*models.Users, error) {

	err := model.DB.Create(user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Searches for the first User related to the given e-mail.
//
// Returns User, Exists, Error.
//
// Ignores Errors like Gorm RecordsNotFound since it is implemented in exists value.
func (model *UserModel) GetByEmail(email string) (User models.Users, Exists bool, Error error) {
	var user models.Users

	err := model.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, false, nil
		} else {
			return user, false, err
		}
	}

	return user, true, nil
}

func (m *UserModel) Delete(u *models.Users) error {
	//
	return nil
}

func (m *UserModel) Update(u *models.Users) error {
	// Update user profile
	return nil
}
