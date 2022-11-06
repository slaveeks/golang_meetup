package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/slaveeks/golang_meetup/src/controllers"
)

const (
	membersUrl = "/members"
	memberUrl  = "/members/:id"
)

type routes struct {
	c controllers.Controller
}

func CreateMemberRoutes(controller controllers.Controller) Route {
	return &routes{
		controller,
	}
}

func (r *routes) Register(e *echo.Echo) {
	fmt.Println("routes created")
	e.GET(membersUrl, r.c.FindAll)
	e.GET(memberUrl, r.c.FindById)
	e.POST(membersUrl, r.c.Create)
	e.DELETE(memberUrl, r.c.Delete)
}
