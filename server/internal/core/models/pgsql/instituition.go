package pgsql

import (
	"errors"

	"gorm.io/gorm"

	"github.com/noeeekr/sch-server/internal/core/models"
)

// database model methods initializer
type InstitutionModel struct {
	DB *gorm.DB
}

func (model *InstitutionModel) FindByUserId(userId uint) (Institution *[]models.Institutions, Exists bool, Error error) {
	var institutions *[]models.Institutions

	if err := model.DB.Table("institutions").
		Select("*").
		Joins("JOIN users_institutions ON users_institutions.institutions_id = institutions.id").
		Where("users_institutions.users_id = ?", userId).
		Scan(&institutions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return institutions, true, nil
}

func (model *InstitutionModel) GetClasses(email string) (User *models.Institutions, Exists bool, Error error) {
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

func (model *InstitutionModel) GetUsers(user *models.Institutions) (*models.Institutions, error) {
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
