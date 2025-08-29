package models

import "github.com/Noeeekr/borm"

// var institutionsTableRequests = &InstitutionsRequests{
// 	Create: transactions.NewRequest(fmt.Sprintf(`
// 		CREATE TABLE IF NOT EXISTS %s (
// 			%s
// 			%s
// 			name VARCHAR(256) NOT NULL
// 		);
// 	`, InstitutionsTableName, DefaultFieldsQuery, SerialId)),
// 	SelectOneById: transactions.NewRequest(fmt.Sprintf(`
// 		SELECT created_at, updated_at, deleted_at, name, id FROM %s WHERE id = $1;
// 	`, InstitutionsTableName)),
// 	SelectOneByName: transactions.NewRequest(fmt.Sprintf(`
// 		SELECT created_at, updated_at, deleted_at, name, id FROM %s WHERE name = $1;
// 	`, InstitutionsTableName)),
// 	InsertMany: transactions.NewRequest(fmt.Sprintf(`
// 		INSERT INTO %s (created_at, updated_at, name)
// 		VALUES $$$$$
// 		RETURNING id;
// 	`, InstitutionsTableName)).AllowValueRepeat("$$$$$", 3),
// 	Drop: transactions.NewRequest(fmt.Sprintf(`
// 		DROP TABLE IF EXISTS %s CASCADE;
// 	`, InstitutionsTableName)),
// 	DeleteOneById: transactions.NewRequest(fmt.Sprintf(`
// 		DELETE FROM %s WHERE id = $1;
// 	`, InstitutionsTableName)),
// 	DeleteOneByName: transactions.NewRequest(fmt.Sprintf(`
// 		DELETE FROM %s WHERE name = $1;
// 	`, InstitutionsTableName)),
// }

type CreateInstitutions struct {
	Name     string `binding:"required" borm:"(CONSTRAINTS, NOT NULL)" json:"name"`
	Email    string `binding:"required" json:"email" borm:"(IGNORE)"`
	Password string `binding:"required" json:"password" borm:"(IGNORE)"`
}

type Institutions struct {
	ID
	DefaultFields
	CreateInstitutions
}

var TableInstitutions *borm.TableRegistry = EnvironmentDatabase.RegisterTable(Institutions{})
