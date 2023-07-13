package endpoint

import (
	"CRUD_Go/internal/app/model"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	Takeuser(name string) (model.User, error)
	Takeusers() ([]model.User, error)
	Adduser(u model.User) (string, error)
	Deleteuser(u model.User) (string, error)
	Updateuser(u model.User) (string, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Hello(ctx echo.Context) error {
	s := fmt.Sprintf("Приветствую в моём приложении\n В нём есть такие методы \n ")
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return errors.New("Error CTX Server")
	}
	return nil
}

func (e *Endpoint) UserTake(ctx echo.Context) error {
	user := model.User{}

	err := ctx.Bind(&user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	user, err = e.s.Takeuser(user.Name)
	return ctx.JSON(http.StatusOK, user)
}

func (e *Endpoint) UserSTake(ctx echo.Context) error {
	users, err := e.s.Takeusers()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, users)
}

func (e *Endpoint) UserAdd(ctx echo.Context) error {
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	ansver, err := e.s.Adduser(user)
	return ctx.JSON(http.StatusOK, ansver)
}

func (e *Endpoint) UserUpdate(ctx echo.Context) error {
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	ansver, err := e.s.Updateuser(user)
	return ctx.JSON(http.StatusOK, ansver)
}

func (e *Endpoint) UserDelete(ctx echo.Context) error {
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	ansver, err := e.s.Deleteuser(user)
	return ctx.JSON(http.StatusOK, ansver)
}
