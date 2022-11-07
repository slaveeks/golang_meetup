package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/slaveeks/golang_meetup/pkg/client"
	"github.com/slaveeks/golang_meetup/src/controllers"
	"github.com/slaveeks/golang_meetup/src/models"
	"github.com/slaveeks/golang_meetup/src/routes"
)

func main() {
	e := echo.New()

	database, _ := client.CreateClient("localhost", "27017", "meetup")

	memberModel := models.CreateMemberModel("members", database, context.TODO())

	memberController := controllers.CreateMemberController(memberModel)

	memberRoutes := routes.CreateMemberRoutes(memberController)
	memberRoutes.Register(e)

	err := e.Start(":1234")

	if err != nil {
		panic(fmt.Sprintf("Erorr while starting a server: %v", err))
	}
}
