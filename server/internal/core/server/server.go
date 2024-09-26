package server

import (
	"net/http"

	configs "github.com/noeeekr/sch-server/config"
	logs "github.com/noeeekr/sch-server/internal/core/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ServeAndListen() error {
	cfg, err := configs.GetConfig()
	if err != nil {
		return err
	}

	var port string = cfg.Addr

	logs.LogInfo.Println("Connecting to database..")

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
	defer sql.Close()

	logs.LogInfo.Println("Starting server..")

	router, err := getRouter(db)
	if err != nil {
		return err
	}

	server := http.Server{
		Addr:     ":" + port,
		Handler:  router,
		ErrorLog: logs.LogErr,
	}

	logs.LogInfo.Printf("\nServer is running on http://localhost:%s", port)

	err = server.ListenAndServe()

	return err
}
