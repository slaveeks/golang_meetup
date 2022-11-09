package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/slaveeks/golang_meetup/src/models"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type memberController struct {
	model *models.MemberModel
}

func (m *memberController) Create(ctx echo.Context) error {
	name := ctx.QueryParam("name")
	email := ctx.QueryParam("email")
	res, err := m.model.Create(name, email)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (m *memberController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	res, err := m.model.Delete(id)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ctx.JSON(http.StatusNotFound, err.Error())
		}

		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (m *memberController) FindAll(ctx echo.Context) error {
	res, err := m.model.FindAll()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (m *memberController) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	res, err := m.model.FindById(id)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ctx.JSON(http.StatusNotFound, err.Error())
		}

		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func CreateMemberController(model *models.MemberModel) Controller {

	return &memberController{model}
}
