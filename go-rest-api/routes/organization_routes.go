package routes

import (
	"go-rest-api/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterOrganizationRoutes(e *echo.Echo) {
	orgs := e.Group("/organizations")

	orgs.POST("", controllers.CreateOrganization)
	orgs.GET("", controllers.ListOrganizations)
	orgs.GET("/:id", controllers.GetOrganizationByID)
	orgs.PUT("/:id", controllers.UpdateOrganization)
	orgs.DELETE("/:id", controllers.DeleteOrganization)
}
