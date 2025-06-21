package models

import (
	"time"
)

type UserRole string

// Define constants for UserRole
const (
	Student    UserRole = "student"
	Teacher    UserRole = "teacher"
	Supervisor UserRole = "supervisor"
	Admin      UserRole = "admin"
)

type Signin struct { // FOR JSON
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=2"`
}

type CommonDbFields struct {
	ID        uint       `gorm:"primary_key;autoincrement:true;unique" json:"id" binding:"required"`
	CreatedAt time.Time  `json:"created_at" binding:"required"`
	UpdatedAt time.Time  `json:"updated_at" binding:"required"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CreateUsers struct { // FOR JSON
	Name     string   `gorm:"not null;size:255" json:"name" binding:"required,min=2,max=255"`
	Email    string   `gorm:"uniqueIndex;size:255" json:"email" binding:"required,email"`
	Password string   `gorm:"not null" json:"password" binding:"required,min=2"`
	Role     UserRole `gorm:"type:user_role;not null;" json:"role" binding:"required"` // Specify the PostgreSQL enum type

	InstitutionId uint `gorm:"not null;index" json:"institution_id" binding:"required"`
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
