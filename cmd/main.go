package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MeiramSh/dostap/internal/config"
	"github.com/MeiramSh/dostap/internal/controller"
	"github.com/MeiramSh/dostap/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg := config.NewConfig()

	db, err := sql.Open("postgres", cfg.Postgres.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := migrate.Migrate(db); err != nil {
		log.Fatal(err)
	}

	app, err := pkg.App()

	if err != nil {
		log.Fatal(err)
	}
	defer app.CloseDBConnection()

	ginRouter := gin.Default()

	controller.Setup(app, ginRouter)

	ginRouter.Run(fmt.Sprintf(":%s", 1136))
}
