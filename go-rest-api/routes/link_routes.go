package routes

import (
	"go-rest-api/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterLinkRoutes(e *echo.Echo) {
	e.POST("/links", controllers.CreateLinkHandler)
	e.GET("/links/:short_code", controllers.GetLinkHandler)
	e.GET("/:short_code", controllers.GetLinkHandler)
	e.DELETE("/links/:short_code", controllers.DeleteLinkHandler)
	e.GET("/links", controllers.SearchLinksHandler)
	e.PUT("/links/:id", controllers.UpdateLinkHandler)
	e.PATCH("/links/:id", controllers.PatchLinkHandler)
}
