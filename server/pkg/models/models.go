package models

import (
	"time"
)

type UserRole string

// Define constants for UserRole
const (
	Student     UserRole = "student"
	Teacher     UserRole = "teacher"
	Institution UserRole = "institution"
)

type Signin struct { // FOR JSON
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=2"`
}

type CommonDbFields struct {
	ID        uint       `gorm:"primary_key" json:"id" binding:"required"`
	CreatedAt time.Time  `json:"created_at" binding:"required"`
	UpdatedAt time.Time  `json:"updated_at" binding:"required"`
	DeletedAt *time.Time `json:"deleted_at" binding:"required"`
}

type CreateUsers struct { // FOR JSON
	Name     string   `gorm:"not null;size:255" json:"name" binding:"required,min=2,max=255"`
	Email    string   `gorm:"uniqueIndex;size:255" json:"email" binding:"required,email"`
	Password string   `gorm:"not null" json:"password" binding:"required,min=2"`
	Role     UserRole `gorm:"type:user_role;not null;default:'student'" json:"role" binding:"required"` // Specify the PostgreSQL enum type

	InstitutionId uint `gorm:"not null;index" json:"institution_id" binding:"required"`
}

type Users struct {
	CommonDbFields
	CreateUsers
	Institution Institutions `gorm:"foreignKey:InstitutionId" json:"-"`

	ProfileImgUrl string `gorm:"default:'./assets/defaultphp.jpg'" json:"profile_img_url" binding:"required"`

	Notifications []Notifications `gorm:"many2many:user_notifications" json:"-"`
	Classes       []Classes       `gorm:"many2many:user_classes;" json:"-"` // Many-to-many relationship with classes
}

type CreateInstitutions struct { // FOR JSON
	Name     string   `gorm:"not null; size:255" json:"name" binding:"required"`
	Email    string   `gorm:"size:255; not null; uniqueIndex" json:"email" binding:"required"`
	Password string   `gorm:"not null" json:"password" binding:"required"`
	Role     UserRole `gorm:"type:user_role;not null;default:'institution'" json:"role" binding:"required"`
}

type Institutions struct {
	CommonDbFields
	CreateInstitutions

	Name          string `gorm:"size:255; not null" json:"name" binding:"required"`
	ProfileImgUrl string `gorm:"default:'./assets/defaultphp.jpg'" json:"profile_img_url" binding:"required"`

	Classes      []Classes `gorm:"foreignKey:InstitutionId" json:"-"`
	Participants []Users   `gorm:"foreignKey:InstitutionId" json:"-"` // This will hold both teachers and students
}

type CreateClasses struct {
	TeacherId uint  `gorm:"not null;index" json:"teacher_id"` // Corrected JSON key
	Teacher   Users `gorm:"foreignKey:TeacherId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"teacher"`

	InstitutionId uint         `gorm:"not null;index" json:"institution_id" binding:"required"` // Foreign key to institution
	Institution   Institutions `gorm:"foreignKey:InstitutionId" json:"-"`
}

// Class Model (Many Students, One Teacher)
type Classes struct {
	CommonDbFields
	CreateClasses

	Name string `gorm:"size:255" json:"name"`

	// Store student IDs directly instead of a slice of Users to avoid recursion
	Students      []Users         `gorm:"many2many:user_classes;" json:"students"` // Store users directly linked to this class
	Notifications []Notifications `gorm:"foreignKey:ClassId" json:"notifications"`
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

	Students []Users `gorm:"many2many:user_notifications"` // Array of student IDs (foreign keys)
}

type UserNotifications struct {
	UserID         uint `gorm:"primaryKey"`
	NotificationId uint `gorm:"primaryKey"`
}

type UserClasses struct {
	UserID  uint `gorm:"primaryKey"`
	ClassID uint `gorm:"primaryKey"`
}
