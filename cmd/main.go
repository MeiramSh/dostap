package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/MeiramSh/dostap/internal/config"
	"github.com/MeiramSh/dostap/internal/controller"
	"github.com/MeiramSh/dostap/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg := config.NewConfig()

	app, err := pkg.App(cfg.Postgres.DSN())
	if err != nil {
		log.Fatal(err)
	}

	err = pkg.Migrate(app.DB)

	if err != nil {
		log.Fatal(err)
	}
	defer app.CloseDBConnection()

	ginRouter := gin.Default()

	controller.Setup(app, ginRouter)

	ginRouter.Run(fmt.Sprintf(":%s", strconv.Itoa(cfg.Port)))

}
