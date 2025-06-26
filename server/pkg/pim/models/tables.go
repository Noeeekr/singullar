package models

import (
	"fmt"
	"time"
)

type TableInfo struct {
	Dependencies *TableDepencies
	query        string
	name         TableName
}

func (t *TableInfo) TableName() TableName {
	return t.name
}
func (t *TableInfo) TableQuery() string {
	return t.query
}

type TableDepencies struct {
	Types  []TypeName
	Tables []TableName
}

type TableName string

const (
	usersTableName         TableName = "users"         // DONE
	institutionsTableName  TableName = "institutions"  // PARTIAL
	classesTableName       TableName = "classes"       // NOT DONE
	notificationsTableName TableName = "notifications" // NOT DONE

	usersInstitutionsTableName  TableName = "usersInstitutions"  // NOT DONE
	usersClassesTableName       TableName = "usersClasses"       // NOT DONE
	usersNotificationsTableName TableName = "usersNotifications" // NOT DONE
)

type DefaultFields struct {
	Id        int        `json:"id" binding:"required"`
	CreatedAt time.Time  `json:"created_at" binding:"required"`
	UpdatedAt time.Time  `json:"updated_at" binding:"required"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TABLE DEFAUT FIELDS
const DefaultFieldsQuery = `
	id         SERIAL      PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ,
`

// TABLE USERS
var UsersTable = TableInfo{
	name: usersTableName,
	query: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			name             VARCHAR(256)   NOT NULL,
			email            VARCHAR(256)   NOT NULL UNIQUE,
			password 	     VARCHAR(256)   NOT NULL,
			institution_id  INT      	    NOT NULL,
			role             %s             NOT NULL,
			profile_picture  VARCHAR(256)   DEFAULT userprofilepicture.jpg,

			FOREIGN KEY (institution_id) REFERENCES %s(id)
		);
	`, usersTableName, DefaultFieldsQuery, UserRoleName, institutionsTableName),
	Dependencies: &TableDepencies{
		Types:  []TypeName{UserRoleName},
		Tables: []TableName{institutionsTableName},
	},
}

type Users struct {
	DefaultFields
	Name           string   `json:"name" binding:"required,min=2,max=255"`
	Email          string   `json:"email" binding:"required,email"`
	Password       string   `json:"password" binding:"required,min=6"`
	ProfilePicture string   `json:"profile_picture"`
	Role           UserRole `json:"role" binding:"required"`

	InstitutionId int `json:"institution_id" binding:"required"`
}

var InstitutionsTable = TableInfo{
	name: institutionsTableName,
	query: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			name VARCHAR(256) NOT NULL
		);
	`, institutionsTableName, DefaultFieldsQuery),
	Dependencies: &TableDepencies{
		Types:  []TypeName{},
		Tables: []TableName{},
	},
}

type Institutions struct {
	DefaultFields
	Name string `json:"name" binding:"required"`
}

var ClassesTable = TableInfo{
	name: classesTableName,
	query: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			name VARCHAR(256) NOT NULL,
			segment VARCHAR(256) NOT NULL,
			series VARCHAR(256) NOT NULL,

			institution_id INT NOT NULL,

			FOREIGN KEY institution_id REFERENCES %s(id)
		);
	`, classesTableName, DefaultFieldsQuery, institutionsTableName),
	Dependencies: &TableDepencies{
		Types:  []TypeName{},
		Tables: []TableName{institutionsTableName},
	},
}

type Classes struct {
	DefaultFields
	Name    string `json:"name" binding:"required"`
	Segment string `json:"segment" binding:"required"`
	Series  string `json:"series" binding:"required"`

	InstitutionId int `json:"institution_id" binding:"required"`
}

var NotificationsTable = TableInfo{
	name: notificationsTableName,
	query: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s
			title		 VARCHAR(256) NOT NULL,
			description  VARCHAR(256) NOT NULL,
			target_id 	 INT 		  NOT NULL,
			target_type  %s 		  NOT NULL
		);
	`, notificationsTableName, DefaultFieldsQuery, UserRoleName),
	Dependencies: &TableDepencies{
		Types: []TypeName{UserRoleName},
		// Classes and institutions since target id may point to one
		Tables: []TableName{classesTableName, institutionsTableName},
	},
}

type Notifications struct {
	DefaultFields

	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	// Notification can be sent to a Class using UserRole=Unkown and TargetId=ClassID
	// Notification can be sent to a Roles using UserRole=Role and TargetId=InstitutionID
	TargetId   int      `json:"target_id" binding:"required"`
	TargetType UserRole `json:"target_type" binding:"required"`
}

var UsersClassesTable = &TableInfo{
	name: usersClassesTableName,
	query: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			user_id INT NOT NULL,
			class_id INT NOT NULL,

			FOREIGN KEY users_id REFERENCES %s(id),
			FOREIGN KEY class_id REFERENCES %s(id)
		);
	`, usersClassesTableName, usersTableName, classesTableName),
	Dependencies: &TableDepencies{
		Types:  []TypeName{},
		Tables: []TableName{usersTableName, classesTableName},
	},
}

type UsersClasses struct {
	UserId        int
	InstitutionId int
}

var UsersNotificationsTable = &TableInfo{
	name: usersNotificationsTableName,
	query: fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			user_id INT NOT NULL,
			notification_id INT NOT NULL,

			FOREIGN KEY users_id REFERENCES %s(id),
			FOREIGN KEY notification_id REFERENCES %s(id)
		);
	`, usersNotificationsTableName, usersTableName, notificationsTableName),
	Dependencies: &TableDepencies{
		Types:  []TypeName{},
		Tables: []TableName{usersTableName, notificationsTableName},
	},
}

type UsersNotifications struct {
	UserId         int
	NotificationId int
}

// Tables
//	- Users
// 		Types:
// 			Migration: DONE
// 			API JSON: DONE
// 			API CREATE: DONE
//  - Institutions
//  		Migration:
//  		API JSON:
//			API CREATE:
//	- Notifications
//  		Migration:
//  		API JSON:
//			API CREATE:
//  - Classes
//  		Migration:
//  		API JSON:
//			API CREATE:
//
// SubTables
//  - UsersNotifications
//		Depencies
//			- Users - Notifications
//
// 	- UsersInstitutions
// 		Depencies
// 			- Institutions - Users
//
//  - UsersClasses
// 		Depencies
// 			- Classes - Users
//
//  Institutions Config
//		(CHECK ACCESS COOKIE) Then find users in UsersInstitutions WHERE userID == user, instID == institution, role == ADMIN access INSTITUION CONFIG PANEL
//	Classes Config
//		(CHECK ACCESS COOKIE) Then find users in classesUsers WHERE userId == user, role == TEACHER
//  Notifications
//  	Same from above
/*

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
*/

/*
DONE:
type Users struct {
	CommonDbFields
	CreateUsers
	// SHould be table  | Institutions []Institutions `gorm:"many2many:users_institutions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	// SHould be table  | Notifications []Notifications `gorm:"many2many:users_notifications" json:"-"`
	// Should be table | Classes       []Classes       `gorm:"many2many:users_classes;" json:"-"` // Many-to-many relationship with classes
}
type UserRole string
// Define constants for UserRole
const (
	Student    UserRole = "student"
	Teacher    UserRole = "teacher"
	Supervisor UserRole = "supervisor"
	Admin      UserRole = "admin"
)

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
type Institutions struct {
	CommonDbFields
	Name          string `gorm:"size:255; not null" json:"name" binding:"required"`
	ProfileImgUrl string `gorm:"default:'./assets/defaultpfp.jpg'" json:"profile_img_url" binding:"required"`

	Classes []Classes `gorm:"foreignKey:InstitutionId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"-"`
	Users   []Users   `gorm:"many2many:users_institutions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"-"` // This will hold both teachers and students
}
*/
