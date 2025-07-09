package migrations

import "github.com/Noeeekr/singullar/server/internal/database/models"

type Context struct {
	alreadyCreatedTables map[models.TableName]bool
	alreadyCreatedTypes  map[models.TypeName]bool
}

type Configuration struct {
	IgnoreExisting   bool
	RecreateExisting bool
}

type RequestCreateDatabase struct {
	Database string
	User     string
}
