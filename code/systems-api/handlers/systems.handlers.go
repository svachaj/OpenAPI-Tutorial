package handlers

import (
	"net/http"
	"panda/apigateway/models"
	"panda/apigateway/services"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type SystemsHandlers struct {
	systemsService services.ISystemsService
}

type ISystemsHandlers interface {
	CreateNewSystem() echo.HandlerFunc
	DeleteSystemByCode() echo.HandlerFunc
	GetSystemByCode() echo.HandlerFunc
	GetSystemsByNameOrCode() echo.HandlerFunc
	GetSystemMaintenance() echo.HandlerFunc
	DeleteConfigurationByKeyAndSystemCode() echo.HandlerFunc
	GetSystemConfigurationBySystemCode() echo.HandlerFunc
	GetSystemTimeValueLogs() echo.HandlerFunc
	RecreateDatabaseData() echo.HandlerFunc
}

// NewCommentsHandlers Comments handlers constructor
func NewSystemsHandlers(systemsSvc services.ISystemsService) ISystemsHandlers {
	return &SystemsHandlers{systemsService: systemsSvc}
}

func (h *SystemsHandlers) CreateNewSystem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var system models.System
		err := c.Bind(&system)
		if err != nil {
			return c.JSON(401, "Invalid system data")
		}
		result, err := h.systemsService.CreateNewSystem(system)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(500, "General server error")
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (h *SystemsHandlers) DeleteSystemByCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		systemCode := c.Param("systemCode")
		result, err := h.systemsService.DeleteSystemByCode(systemCode)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(500, "General server error")
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (h *SystemsHandlers) GetSystemByCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		systemCode := c.Param("systemCode")
		result, err := h.systemsService.GetSystemByCode(systemCode)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(401, "System not found")
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (h *SystemsHandlers) GetSystemsByNameOrCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		searchText := strings.ToLower(c.QueryParam("searchText"))
		var limit int32
		if limit_param, err := strconv.ParseInt(c.QueryParam("limit"), 10, 64); err == nil {
			limit = int32(limit_param)
		} else {
			return c.JSON(401, "Invalid limit")
		}

		result, err := h.systemsService.GetSystemsByNameOrCode(searchText, limit)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(500, "General server error")
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (h *SystemsHandlers) GetSystemMaintenance() echo.HandlerFunc {
	return func(c echo.Context) error {
		systemCode := c.QueryParam("systemCode")
		result, err := h.systemsService.GetSystemMaintenance(systemCode)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(500, "General server error")
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (h *SystemsHandlers) DeleteConfigurationByKeyAndSystemCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		systemCode := c.Param("systemCode")
		key := c.QueryParam("key")
		result, err := h.systemsService.DeleteConfigurationByKeyAndSystemCode(systemCode, key)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(500, "General server error")
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (h *SystemsHandlers) GetSystemConfigurationBySystemCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		systemCode := c.Param("systemCode")
		result, err := h.systemsService.GetSystemConfigurationBySystemCode(systemCode)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(500, "General server error")
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (h *SystemsHandlers) GetSystemTimeValueLogs() echo.HandlerFunc {
	return func(c echo.Context) error {
		systemCode := c.Param("systemCode")

		result, err := h.systemsService.GetSystemTimeValueLogs(systemCode)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(500, "General server error")
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (h *SystemsHandlers) RecreateDatabaseData() echo.HandlerFunc {
	return func(c echo.Context) error {

		result, err := h.systemsService.RecreateDatabaseData()
		if err != nil {
			log.Error(err.Error())
			return c.JSON(500, "General server error")
		}
		return c.JSON(http.StatusOK, result)
	}
}
