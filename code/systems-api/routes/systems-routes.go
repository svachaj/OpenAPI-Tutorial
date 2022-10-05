package routes

import (
	"panda/apigateway/handlers"

	"github.com/labstack/echo/v4"
)

func MapSystemsRoutes(g *echo.Group, h handlers.ISystemsHandlers, jwtMiddleware echo.MiddlewareFunc) {
	// Create new system route
	g.POST("/system", h.CreateNewSystem(), jwtMiddleware)
	g.GET("/systems", h.GetSystemsByNameOrCode())
	g.GET("/system/:systemCode", h.GetSystemByCode())
	g.DELETE("/system/:systemCode", h.DeleteSystemByCode(), jwtMiddleware)

	g.GET("/system/configuration/:systemCode", h.GetSystemConfigurationBySystemCode())
	g.DELETE("/system/configuration/:systemCode", h.DeleteConfigurationByKeyAndSystemCode(), jwtMiddleware)

	g.GET("/system/maintenance", h.GetSystemMaintenance())

	g.GET("/system/time-value-logs/:systemCode", h.GetSystemTimeValueLogs())

	g.POST("/database/deleteAndInitNewData", h.RecreateDatabaseData(), jwtMiddleware)
}
