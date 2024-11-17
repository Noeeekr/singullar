package server

import (
	"net/http"

	configs "github.com/noeeekr/sch-server/config"
	logs "github.com/noeeekr/sch-server/internal/core/log"
	models "github.com/noeeekr/sch-server/internal/core/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ServeAndListen() error {
	// GET VARIABLES
	cfg, err := configs.GetConfig()
	if err != nil {
		return err
	}

	var port string = cfg.Addr

	logs.LogInfo.Println("Connecting to database..")

	// CONNECT & SYNC DB
	db, err := gorm.Open(postgres.Open(
		cfg.DatabaseWR_ConnectString),
		&gorm.Config{
			Logger: logger.Discard,
		},
	)

	if err != nil {
		return err
	}

	sql, err := db.DB()
	if err != nil {
		return err
	}

	err = migrate(db)
	if err != nil {
		logs.LogErr.Fatal(err)
	}

	defer sql.Close()

	logs.LogInfo.Println("Starting server..")

	// SETUP ROUTER

	router, err := getRouter(db)
	if err != nil {
		return err
	}

	// SETUP SERVER

	server := http.Server{
		Addr:     ":" + port,
		Handler:  router,
		ErrorLog: logs.LogErr,
	}

	logs.LogInfo.Printf("\nServer is running on http://localhost:%s", port)

	err = server.ListenAndServe()

	return err
}

func migrate(db *gorm.DB) error {
	if err := db.Exec(`DO $$ 
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
				CREATE TYPE user_role AS ENUM ('student', 'teacher', 'admin', 'supervisor');
			END IF;
		END $$;`).Error; err != nil {
		return err
	}

	err := db.AutoMigrate(
		&models.Users{},
		&models.Institutions{},
		&models.Classes{},
		&models.Notifications{},
	)
	if err != nil {
		return err
	}

	return nil
}
