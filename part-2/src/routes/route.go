package routes

import "github.com/labstack/echo/v4"

type Route interface {
	Register(e *echo.Echo)
}
