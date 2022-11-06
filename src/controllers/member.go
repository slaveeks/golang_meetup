package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/slaveeks/golang_meetup/src/models"
)

type memberController struct {
	model *models.MemberModel
}

func (m *memberController) Create(ctx echo.Context) error {
	fmt.Println("123")
	name := ctx.QueryParam("name")
	email := ctx.QueryParam("email")
	res, _ := m.model.Create(name, email)

	return ctx.JSON(200, res)
}

func (m *memberController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	res, _ := m.model.Delete(id)

	return ctx.JSON(200, res)
}

func (m *memberController) FindAll(ctx echo.Context) error {
	fmt.Println("123")
	res, _ := m.model.FindAll()

	return ctx.JSON(200, res)
}

func (m *memberController) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	res, _ := m.model.FindById(id)

	return ctx.JSON(200, res)
}

func CreateMemberController(model *models.MemberModel) Controller {

	return &memberController{model}
}
