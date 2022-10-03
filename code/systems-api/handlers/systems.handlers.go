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

// // CreateNewSystem godoc
// // @Summary Create new system
// // @Description Create new system and return new System ID
// // @Tags Systems
// // @Accept json
// // @Produce json
// // @Success 200
// // @Router /system [post]
// // @Security ApiKeyAuth
// func (h *SystemsHandlers) CreateNewSystem() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		name := c.FormValue("name")
// 		description := c.FormValue("description")
// 		systemCode := c.FormValue("systemCode")
// 		systemAlias := c.FormValue("systemAlias")
// 		facilityZone := c.FormValue("facilityZone")
// 		location := c.FormValue("location")
// 		ownerPerson := c.FormValue("ownerPerson")
// 		responsiblePerson := c.FormValue("responsiblePerson")
// 		maintainedByPerson := c.FormValue("maintainedByPerson")

// 		system := models.System{
// 			Name:               &name,
// 			Description:        &description,
// 			SystemCode:         &systemCode,
// 			SystemAlias:        &systemAlias,
// 			FacilityZone:       &facilityZone,
// 			Location:           &location,
// 			OwnerPerson:        &ownerPerson,
// 			ResponsiblePerson:  &responsiblePerson,
// 			MaintainedByPerson: &maintainedByPerson,
// 		}

// 		newSystemID, err := h.systemsService.CreateNewSystem(system)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, "Unexpected server error: "+err.Error())
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"id": newSystemID})
// 	}
// }

// func (h *SystemsHandlers) UpdateSystem() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		formParams, _ := c.FormParams()
// 		system := models.System{}
// 		name := c.FormValue("name")
// 		description := c.FormValue("description")
// 		systemCode := c.FormValue("systemCode")
// 		systemAlias := c.FormValue("systemAlias")
// 		facilityZone := c.FormValue("facilityZone")
// 		location := c.FormValue("location")
// 		ownerPerson := c.FormValue("ownerPerson")
// 		responsiblePerson := c.FormValue("responsiblePerson")
// 		maintainedByPerson := c.FormValue("maintainedByPerson")

// 		if formParams.Has("name") {
// 			system.Name = &name
// 		}
// 		if formParams.Has("description") {
// 			system.Description = &description
// 		}
// 		if formParams.Has("systemCode") {
// 			system.SystemCode = &systemCode
// 		}
// 		if formParams.Has("systemAlias") {
// 			system.SystemAlias = &systemAlias
// 		}
// 		if formParams.Has("facilityZone") {
// 			system.FacilityZone = &facilityZone
// 		}
// 		if formParams.Has("location") {
// 			system.Location = &location
// 		}
// 		if formParams.Has("ownerPerson") {
// 			system.OwnerPerson = &ownerPerson
// 		}
// 		if formParams.Has("responsiblePerson") {
// 			system.ResponsiblePerson = &responsiblePerson
// 		}
// 		if formParams.Has("maintainedByPerson") {
// 			system.MaintainedByPerson = &maintainedByPerson
// 		}

// 		if formParams.Has("id") {
// 			if vid, err := strconv.ParseInt(c.FormValue("id"), 10, 64); err == nil {
// 				system.Id = vid
// 			}
// 		}

// 		msg, err := h.systemsService.UpdateSystem(system)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, "Unexpected server error: "+err.Error())
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"Result": msg})
// 	}
// }

// // CreateNewSubSystem godoc
// // @Summary Create new subsystem
// // @Description Create new subsystem as child of and existing System. You can pass existing System id, uid or name(be careful because the name could not be unique)
// // @Tags Systems
// // @Accept json
// // @Produce json
// // @Success 200
// // @Router /system/subsystem [post]
// // @Security ApiKeyAuth
// func (h *SystemsHandlers) CreateNewSubsystem() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		name := c.FormValue("name")
// 		description := c.FormValue("description")
// 		systemCode := c.FormValue("systemCode")
// 		systemAlias := c.FormValue("systemAlias")
// 		facilityZone := c.FormValue("facilityZone")
// 		location := c.FormValue("location")
// 		ownerPerson := c.FormValue("ownerPerson")
// 		responsiblePerson := c.FormValue("responsiblePerson")
// 		maintainedByPerson := c.FormValue("maintainedByPerson")

// 		system := models.System{
// 			Name:               &name,
// 			Description:        &description,
// 			SystemCode:         &systemCode,
// 			SystemAlias:        &systemAlias,
// 			FacilityZone:       &facilityZone,
// 			Location:           &location,
// 			OwnerPerson:        &ownerPerson,
// 			ResponsiblePerson:  &responsiblePerson,
// 			MaintainedByPerson: &maintainedByPerson,
// 		}

// 		var parentName string
// 		var parentId int64 = -1
// 		var parentUid string
// 		formParams, _ := c.FormParams()

// 		if formParams.Has("parentName") {
// 			parentName = c.FormValue("parentName")
// 		}
// 		if formParams.Has("parentUid") {
// 			parentUid = c.FormValue("parentUid")
// 		}
// 		if formParams.Has("parentId") {
// 			if vid, err := strconv.ParseInt(c.FormValue("parentId"), 10, 64); err == nil {
// 				parentId = vid
// 			}
// 		}

// 		newSystemID, err := h.systemsService.CreateNewSubsystem(system, parentId, parentUid, parentName)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, "Unexpected server error: "+err.Error())
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"id": newSystemID})
// 	}
// }

// func (h *SystemsHandlers) CreateNewHierarchicalRelationship() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		var parentName string
// 		var parentId int64 = -1
// 		var parentUid string
// 		var childName string
// 		var childId int64 = -1
// 		var childUid string

// 		formParams, _ := c.FormParams()

// 		if formParams.Has("parentName") {
// 			parentName = c.FormValue("parentName")
// 		}
// 		if formParams.Has("parentUid") {
// 			parentUid = c.FormValue("parentUid")
// 		}
// 		if formParams.Has("parentId") {
// 			if vid, err := strconv.ParseInt(c.FormValue("parentId"), 10, 64); err == nil {
// 				parentId = vid
// 			}
// 		}
// 		if formParams.Has("childName") {
// 			childName = c.FormValue("childName")
// 		}
// 		if formParams.Has("childUid") {
// 			childUid = c.FormValue("childUid")
// 		}
// 		if formParams.Has("childId") {
// 			if vid, err := strconv.ParseInt(c.FormValue("childId"), 10, 64); err == nil {
// 				childId = vid
// 			}
// 		}

// 		newSystemID, err := h.systemsService.CreateParentChildRelationship(parentId, parentUid, parentName, childId, childUid, childName)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, "Unexpected server error: "+err.Error())
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"id": newSystemID})
// 	}
// }

// // Delete System and all its relationships
// func (h *SystemsHandlers) DeleteSystemAndRelationships() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		formParams, _ := c.FormParams()

// 		var systemId int64
// 		if formParams.Has("systemId") {
// 			if vid, err := strconv.ParseInt(c.FormValue("systemId"), 10, 64); err == nil {
// 				systemId = vid
// 			}
// 		}

// 		msg, err := h.systemsService.DeleteSystemAndRelationships(systemId)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, "Unexpected server error: "+err.Error())
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"Result": msg})
// 	}
// }

// // Delete relationship by parent and child ids
// func (h *SystemsHandlers) DeleteRelationshipByParentChildIds() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		formParams, _ := c.FormParams()

// 		var parentId int64
// 		if formParams.Has("parentId") {
// 			if vid, err := strconv.ParseInt(c.FormValue("parentId"), 10, 64); err == nil {
// 				parentId = vid
// 			}
// 		}
// 		var childId int64
// 		if formParams.Has("childId") {
// 			if vid, err := strconv.ParseInt(c.FormValue("childId"), 10, 64); err == nil {
// 				childId = vid
// 			}
// 		}

// 		msg, err := h.systemsService.DeleteRelationshipByParentChildIds(parentId, childId)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, "Unexpected server error: "+err.Error())
// 		}

// 		return c.JSON(http.StatusOK, echo.Map{"Result": msg})
// 	}
// }
