package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/locatemybeer/lmb-back/apis"
	"gitlab.com/locatemybeer/lmb-back/app"
	"gitlab.com/locatemybeer/lmb-back/daos"
	"gitlab.com/locatemybeer/lmb-back/errors"
	"gitlab.com/locatemybeer/lmb-back/services"
)

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}



	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	//connecting to DB
	db, err := gorm.Open("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SetLogger(logger)
	db.Begin()
	// wire up API routing
	http.Handle("/", buildRouter(logger,db))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger,db *gorm.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort()  // skip all other middlewares/handlers
		return c.Write("OK " + app.Version)
	})

	router.Use(
		app.Init(logger, db),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
	)

	rg := router.Group("/v1")

	rg.Post("/auth", apis.Auth(app.Config.JWTSigningKey))
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  apis.JWTHandler,
	}))

	beerDAO := daos.NewBeerDAO()
	apis.ServeBeerResource(rg, services.NewBeerService(beerDAO))

	barDAO := daos.NewBarDAO()
	apis.ServeBarResource(rg, services.NewBarService(barDAO))
	// wire up more resource APIs here

	return router
}