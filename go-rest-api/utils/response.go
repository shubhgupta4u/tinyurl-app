package utils

import "github.com/labstack/echo/v4"

func SuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}

func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, map[string]interface{}{
		"status":  "error",
		"message": message,
	})
}
