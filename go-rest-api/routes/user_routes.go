package routes

import (
	"go-rest-api/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo) {
	users := e.Group("/users")

	users.POST("", controllers.CreateUser)
	users.GET("/:id", controllers.GetUserByID)
	users.GET("/org/:org_id", controllers.ListUsersByOrg)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
}
