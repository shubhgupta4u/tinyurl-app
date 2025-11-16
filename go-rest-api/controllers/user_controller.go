package controllers

import (
	logger "go-rest-api/middlewares"
	"go-rest-api/models"
	repository "go-rest-api/repositories"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Create User
func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		logger.Logger.Errorf("Failed to bind user: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	if err := repository.CreateUser(&user); err != nil {
		logger.Logger.Errorf("Failed to create user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	logger.Logger.Infof("User created: %s", user.ID)
	createdUser, err := repository.GetUserByID(user.ID)
	if err != nil {
		logger.Logger.Errorf("Failed to get created user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}
	return c.JSON(http.StatusCreated, createdUser)
}

// Get User by ID
func GetUserByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Logger.Errorf("Invalid UUID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "invalid ID"})
	}

	user, err := repository.GetUserByID(id)
	if err != nil {
		logger.Logger.Errorf("Failed to get user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"status": "error", "message": "user not found"})
	}

	return c.JSON(http.StatusOK, user)
}

// List Users by Org
func ListUsersByOrg(c echo.Context) error {
	orgIDStr := c.Param("org_id")
	orgID, err := uuid.Parse(orgIDStr)
	if err != nil {
		logger.Logger.Errorf("Invalid Org UUID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "invalid org ID"})
	}

	users, err := repository.ListUsersByOrg(orgID)
	if err != nil {
		logger.Logger.Errorf("Failed to list users: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

// Update User
func UpdateUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Logger.Errorf("Invalid UUID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "invalid ID"})
	}

	user, err := repository.GetUserByID(id)
	if err != nil {
		logger.Logger.Errorf("Failed to get user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"status": "error", "message": "user not found"})
	}
	var userModel models.User
	if err := c.Bind(&userModel); err != nil {
		logger.Logger.Errorf("Failed to bind user: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	user.Email = userModel.Email
	user.Mobile = userModel.Mobile
	user.Name = userModel.Name
	user.IsActive = userModel.IsActive

	if err := repository.UpdateUser(user); err != nil {
		logger.Logger.Errorf("Failed to update user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	logger.Logger.Infof("User updated: %s", user.ID)
	return c.JSON(http.StatusOK, user)
}

// Delete User
func DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Logger.Errorf("Invalid UUID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "invalid ID"})
	}

	if err := repository.DeleteUserByID(id); err != nil {
		logger.Logger.Errorf("Failed to delete user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	logger.Logger.Infof("User deleted: %s", id)
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "message": "user deleted"})
}
