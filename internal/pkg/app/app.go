package app

import (
	"CRUD_Go/internal/app/endpoint"
	"CRUD_Go/internal/app/service"
	"CRUD_Go/internal/app/storage/db"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

type App struct {
	database *db.Database
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	dsn := os.Getenv("DSN")

	//инициализации всех параметров через New
	app.database, _ = db.New(dsn)
	app.service = service.New(app.database)
	app.endpoint = endpoint.New(app.service)
	app.echo = echo.New()

	app.echo.GET("/", app.endpoint.Hello)
	app.echo.GET("/user", app.endpoint.UserTake)
	app.echo.POST("/user", app.endpoint.UserAdd)
	app.echo.PUT("/user", app.endpoint.UserUpdate)
	app.echo.DELETE("/user", app.endpoint.UserDelete)
	app.echo.GET("/users", app.endpoint.UserSTake)
	return app, nil
}
func (a *App) Run() error {
	fmt.Println("Server Runnig")

	err := a.echo.Start(":8090")
	if err != nil {
		log.Println(errors.New("Error Start Service"))
	}
	return nil
}
