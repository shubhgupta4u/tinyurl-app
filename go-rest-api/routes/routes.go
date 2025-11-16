package routes

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	// Add more route files here
	RegisterLinkRoutes(e)
	RegisterOrganizationRoutes(e)
	RegisterUserRoutes(e)
}
