package models

import "github.com/Noeeekr/borm"

type CreateUsers struct {
	Name          string   `json:"name" binding:"required,min=2,max=255" borm:"(CONSTRAINTS, NOT NULL)"`
	Email         string   `json:"email" binding:"required,email" borm:"(CONSTRAINTS, UNIQUE, NOT NULL)"`
	Password      string   `json:"password" binding:"required,min=6" borm:"(CONSTRAINTS, NOT NULL)"`
	InstitutionId int      `json:"institution_id" binding:"required" borm:"(NAME, institution_id) (CONSTRAINTS, NOT NULL) (FOREIGN KEY, institutions, id)"`
	Role          UserRole `json:"role" binding:"required" borm:"(TYPE, user_roles) (CONSTRAINTS, NOT NULL)"`

	// Perhaps refactor into a StudentUser table if grows
	Segment *UserSegments `json:"segment" binding:"required" borm:"(TYPE, user_segments)"`
}

type Users struct {
	ID
	DefaultFields
	CreateUsers

	ProfilePicture string `json:"profile_picture" borm:"(NAME, profile_picture) (CONSTRAINTS, DEFAULT 'default_user_pfp')"`
}

func CreateUser(name, email, password string, institutionId int, role UserRole, segment *UserSegments) *CreateUsers {
	return &CreateUsers{
		Name:          name,
		Email:         email,
		Password:      password,
		InstitutionId: institutionId,
		Role:          role,
		Segment:       segment,
	}
}

var TableUsers *borm.TableRegistry = EnvironmentDatabase.RegisterTable(Users{}).
	NeedRoles(TypeUserRole, TypeUserSegment).
	NeedTables(TableInstitutions)

	// 	ProfilePicture string `json:"profile_picture"`
	// }

	// var usersTableRequests = &UsersRequests{
	// 	Create: transactions.NewRequest(fmt.Sprintf(`
	// 		CREATE TABLE IF NOT EXISTS %s (
	// 			%s
	// 			%s
	// 			name             VARCHAR(256)   NOT NULL,
	// 			email            VARCHAR(256)   NOT NULL UNIQUE,
	// 			password 	     VARCHAR(256)   NOT NULL,
	// 			institution_id   INT      	    NOT NULL,
	// 			profile_picture  VARCHAR(256)   DEFAULT 'userprofilepicture.jpg',
	// 			role             %s             NOT NULL,
	// 			segment 		 %s,

	// 			CONSTRAINT fk_institutions
	// 				FOREIGN KEY (institution_id)
	// 				REFERENCES %s(id)
	// 				ON DELETE CASCADE
	// 		);
	// 	`, usersTableName, DefaultFieldsQuery, SerialId, UserRolesTypeName, userSegmentTypeName, InstitutionsTableName)),
	// 	InsertMany: transactions.NewRequest(fmt.Sprintf(`
	// 		INSERT INTO %s (created_at, updated_at, name, email, password, institution_id, role, segment)
	// 		VALUES %s
	// 		RETURNING created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture, segment;
	// 	`, usersTableName, placeholder)).AllowValueRepeat(placeholder, 8),
	// 	SelectOneByEmail: transactions.NewRequest(fmt.Sprintf(`
	// 		SELECT created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture, segment
	// 		FROM %s
	// 		WHERE email = $1;
	// 	`, usersTableName)),
	// 	SelectOneById: transactions.NewRequest(fmt.Sprintf(`
	// 		SELECT created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture, segment
	// 		FROM %s
	// 		WHERE id = $1;
	// 	`, usersTableName)),
	// 	SelectManyByInstitutionId: transactions.NewRequest(fmt.Sprintf(`
	// 		SELECT created_at, updated_at, deleted_at, name, email, password, institution_id, role, id, profile_picture, segment
	// 		FROM %s
	// 		WHERE institution_id = $1;
	// 	`, usersTableName)),
	// 	Drop: transactions.NewRequest(fmt.Sprintf(`
	// 		DROP TABLE IF EXISTS %s CASCADE;
	// 	`, usersTableName)),
	// 	DeleteOneByEmail: transactions.NewRequest(fmt.Sprintf(`
	// 		DELETE FROM %s WHERE email = $1;
	// 	`, usersTableName)),
	// 	DeleteOneById: transactions.NewRequest(fmt.Sprintf(`
	// 		DELETE FROM %s WHERE id = $1;
	// 	`, usersTableName)),
	// }
