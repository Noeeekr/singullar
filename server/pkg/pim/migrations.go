package pim

import (
	"database/sql"
	"fmt"
)

// Should be put appended before other fields
var defaultFields = `
ID        INT         PRIMARY KEY,
CreatedAt TIMESTAMPTZ NOT NULL,
UpdatedAt TIMESTAMPTZ NOT NULL,
DeletedAt TIMESTAMPTZ NOT NULL,
`

/*
Managers

	-> Get Query Manager
		-> Method: Close(): Closes the database connection if no other manager is using this connection
	-> Get Migration Manager
		-> Method: Close(): Closes the database connection if no other manager is using this connection
*/
type MigrationsManager struct {
	db *sql.DB
}

func NewMigrationManager(db *sql.DB) *MigrationsManager {
	return &MigrationsManager{
		db: db,
	}
}

type QueryStatus struct {
	Description string
	Status      Status
}

type Status int

const (
	StatusSuccess Status = iota + 999
	StatusFailedTransaction
	StatusFailedTransactionStart
	StatusFailedTransactionRollback
	StatusUnregisteredMigration
)

type TableName string

const (
	USERS TableName = "users"
)

func (m *MigrationsManager) Close() error {
	return m.db.Close()
}
func (m *MigrationsManager) Ping() error {
	return m.db.Ping()
}
func (m *MigrationsManager) Table(table TableName, disableDefaults bool) *QueryStatus {
	switch table {
	case USERS:
		return m.migrateUsers(disableDefaults)
	}

	return &QueryStatus{
		Status:      StatusUnregisteredMigration,
		Description: "Migration Not Registered",
	}
}
func (m *MigrationsManager) DropTable(table TableName) *QueryStatus {
	query := "DROP TABLE IF EXISTS " + string(table)

	tx, err := m.db.Begin()
	if err != nil {
		return &QueryStatus{
			Status:      StatusFailedTransactionStart,
			Description: "Unable to start transaction. " + err.Error(),
		}
	}

	if _, err := tx.Query(query); err != nil {
		if err := tx.Rollback(); err != nil {
			return &QueryStatus{
				Status:      StatusFailedTransactionRollback,
				Description: "Migration transaction failed. Unable to rollback: " + err.Error(),
			}
		}
		return &QueryStatus{
			Status:      StatusFailedTransaction,
			Description: "Migration transaction failed. Rollback executed: " + err.Error(),
		}
	}

	if err := tx.Commit(); err != nil {
		return &QueryStatus{
			Status:      StatusFailedTransaction,
			Description: "Transaction Failed: " + err.Error(),
		}
	}

	return &QueryStatus{
		Status:      StatusSuccess,
		Description: "Migration transaction sucessfull.",
	}
}

func (m *MigrationsManager) migrateUsers(disableDefaults bool) *QueryStatus {
	if disableDefaults {
		defaultFields = ""
	}

	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS users (
			%s
			Name     VARCHAR(256)   NOT NULL,
			Email    VARCHAR(256)   NOT NULL UNIQUE,
			Password VARCHAR(256)   NOT NULL
		)
	`, defaultFields)

	tx, err := m.db.Begin()
	if err != nil {
		return &QueryStatus{
			Status:      StatusFailedTransactionStart,
			Description: "Unable to start transaction. " + err.Error(),
		}
	}

	if _, err := tx.Query(query); err != nil {
		if err := tx.Rollback(); err != nil {
			return &QueryStatus{
				Status:      StatusFailedTransactionRollback,
				Description: "Migration transaction failed. Unable to rollback: " + err.Error(),
			}
		}
		return &QueryStatus{
			Status:      StatusFailedTransaction,
			Description: "Migration transaction failed. Rollback executed: " + err.Error(),
		}
	}

	if err := tx.Commit(); err != nil {
		return &QueryStatus{
			Status:      StatusFailedTransaction,
			Description: "Transaction Failed: " + err.Error(),
		}
	}

	return &QueryStatus{
		Status:      StatusSuccess,
		Description: "Migration transaction sucessfull.",
	}
}

// ======================================================
// ======================================================
// ======================================================
/*

type Signin_ struct { // FOR JSON
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=2"`
}

type CommonDbFields struct {
	ID        uint       `gorm:"primary_key;autoincrement:true;unique" json:"id" binding:"required"`
	CreatedAt time.Time  `json:"created_at" binding:"required"`
	UpdatedAt time.Time  `json:"updated_at" binding:"required"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Users struct {
	CommonDbFields
	CreateUsers
	Institutions []Institutions `gorm:"many2many:users_institutions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`

	ProfileImgUrl string `gorm:"default:'./assets/defaultpfp.jpg'" json:"profile_img_url" binding:"required"`

	Notifications []Notifications `gorm:"many2many:users_notifications" json:"-"`
	Classes       []Classes       `gorm:"many2many:users_classes;" json:"-"` // Many-to-many relationship with classes
}

type Institutions struct {
	CommonDbFields
	Name          string `gorm:"size:255; not null" json:"name" binding:"required"`
	ProfileImgUrl string `gorm:"default:'./assets/defaultpfp.jpg'" json:"profile_img_url" binding:"required"`

	Classes []Classes `gorm:"foreignKey:InstitutionId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"-"`
	Users   []Users   `gorm:"many2many:users_institutions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"-"` // This will hold both teachers and students
}

type CreateClasses struct {
	TeacherId *uint `gorm:"index" json:"teacher_id"`
	Teacher   Users `gorm:"foreignKey:TeacherId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`

	InstitutionId uint         `gorm:"not null;index" json:"institution_id" binding:"required"` // Foreign key to institution
	Institution   Institutions `gorm:"foreignKey:InstitutionId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"-"`

	Name    string `gorm:"size:255" json:"name"`
	Segment string `gorm:"not null" json:"segment"`
	Series  string `gorm:"not null" json:"series"`
}

// Class Model (Many Students, One Teacher)
type Classes struct {
	CommonDbFields
	CreateClasses

	// Store student IDs directly instead of a slice of Users to avoid recursion
	Students      []Users         `gorm:"many2many:users_classes;constraint:OnDelete:CASCADE" json:"students"` // Store users directly linked to this class
	Notifications []Notifications `gorm:"foreignKey:ClassId;constraint:OnDelete:CASCADE" json:"notifications"`
}

type CreateNotifications struct {
	TeacherId uint  `gorm:"not null;index" json:"teacher_id" binding:"required"`
	Teacher   Users `gorm:"foreignKey:TeacherId"`

	ClassId uint    `gorm:"not null;index" json:"class_id" binding:"required"` // Foreign key to the class
	Class   Classes `gorm:"foreignKey:ClassId"`
}

// Notification Model (One Teacher, Many Students)
type Notifications struct {
	CommonDbFields
	CreateNotifications

	Students []Users `gorm:"many2many:users_notifications"` // Array of student IDs (foreign keys)
}

// OTHER MODELS

type CreateStudent struct {
	Name  string `json:"name" binding:"required"`
	ID    uint   `json:"id" binding:"required"`
	Email string `json:"email" binding:"required"`
}

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
*/
