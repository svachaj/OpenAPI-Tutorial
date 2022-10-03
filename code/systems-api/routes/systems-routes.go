package routes

import (
	"panda/apigateway/handlers"

	"github.com/labstack/echo/v4"
)

func MapSystemsRoutes(g *echo.Group, h handlers.ISystemsHandlers, jwtMiddleware echo.MiddlewareFunc) {
	// Create new system route
	g.POST("/system", h.CreateNewSystem(), jwtMiddleware)
	g.GET("/systems", h.GetSystemsByNameOrCode())
}
