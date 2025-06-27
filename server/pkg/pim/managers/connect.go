package managers

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Noeeekr/singullar/server/pkg/common"
	"github.com/Noeeekr/singullar/server/pkg/common/configs"
	_ "github.com/lib/pq"
)

func GetConnectionStringFromFiles(args ...string) string {
	for i, arg := range args {
		os.Args[i+1] = arg
	}

	root_path := common.GetExecutableDir()
	if root_path == "" {
		panic("Failed to get executable path")
	}

	relative_path := common.GetFirstArgument()
	if relative_path == "" {
		panic("Please provide the path to the environment file.")
	}

	env_path := filepath.Join(root_path, relative_path)
	if err := configs.Parse(env_path); err != nil {
		panic("Status: Failed to parse environment file: " + err.Error())
	}
	var env PostgrestManagerEnvironment

	if err := configs.Scan(&env); err != nil {
		panic("Status: Failed to scan environment file into object: " + err.Error())
	}

	return ParseConnectionString(
		env.POSTGRES_CONTAINER_NAME,
		env.POSTGRES_USER,
		env.POSTGRES_USER_PASSWORD,
	)
}

func ParseConnectionString(host, user, passwd string) string {
	return fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		host, user, passwd, user,
	)
}

func Connect(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, errors.New("Unable to connect to postgres: " + err.Error())
	}

	return db, nil
}
