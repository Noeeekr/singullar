package pgsql

import (
	"gorm.io/gorm"

	"github.com/noeeekr/sch-server/pkg/models"
)

// database model methods initializer
type InstitutionModel struct {
	DB *gorm.DB
}

func (model *InstitutionModel) GetByEmail(email string) (User *models.Institutions, Exists bool, Error error) {
	// var _institutions models.Institutions
	// err := model.DB.Model(&models.Institutions{}).Where("email = ?", email).First(&_institutions).Error
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return _institutions, false, nil
	// 	} else {
	// 		return _institutions, false, err
	// 	}
	// }
	// return _institutions, true, nil
	return nil, false, nil
}

func (model *InstitutionModel) Insert(user *models.Institutions) (*models.Institutions, error) {
	// err := model.DB.Model(&models.Institutions{}).Create(user).Error
	// if err != nil {
	// 	return user, err
	// }
	// return user, nil
	return nil, nil
}

// Searches for the first User related to the given e-mail.
//
// Returns User, Exists, Error.
//
// Ignores Errors like Gorm RecordsNotFound since it is implemented in exists value.

func (m *InstitutionModel) Delete(u *models.Users) error {
	//
	return nil
}

func (m *InstitutionModel) Update(u *models.Users) error {
	// Update user profile
	return nil
}
