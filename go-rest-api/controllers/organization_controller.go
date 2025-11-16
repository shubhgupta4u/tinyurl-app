package controllers

import (
	logger "go-rest-api/middlewares"
	"go-rest-api/models"
	repository "go-rest-api/repositories"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Create Organization
func CreateOrganization(c echo.Context) error {
	var org models.Organization
	if err := c.Bind(&org); err != nil {
		logger.Logger.Errorf("Failed to bind organization: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	if err := repository.CreateOrganization(&org); err != nil {
		logger.Logger.Errorf("Failed to create organization: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	logger.Logger.Infof("Organization created: %s", org.ID)
	return c.JSON(http.StatusCreated, org)
}

// Get Organization by ID
func GetOrganizationByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Logger.Errorf("Invalid UUID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "invalid ID"})
	}

	org, err := repository.GetOrganizationByID(id)
	if err != nil {
		logger.Logger.Errorf("Failed to get organization: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}
	if org == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"status": "error", "message": "organization not found"})
	}

	return c.JSON(http.StatusOK, org)
}

// List Organizations
func ListOrganizations(c echo.Context) error {
	orgs, err := repository.ListOrganizations()
	if err != nil {
		logger.Logger.Errorf("Failed to list organizations: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, orgs)
}

// Update Organization
func UpdateOrganization(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Logger.Errorf("Invalid UUID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "invalid ID"})
	}

	org, err := repository.GetOrganizationByID(id)
	if err != nil {
		logger.Logger.Errorf("Failed to get organization: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}
	if org == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"status": "error", "message": "organization not found"})
	}
	var orgModel models.Organization
	if err := c.Bind(&orgModel); err != nil {
		logger.Logger.Errorf("Failed to bind organization: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": err.Error()})
	}
	org.Name = orgModel.Name
	if err := repository.UpdateOrganization(org); err != nil {
		logger.Logger.Errorf("Failed to update organization: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	logger.Logger.Infof("Organization updated: %s", org.ID)
	return c.JSON(http.StatusOK, org)
}

// Delete Organization
func DeleteOrganization(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Logger.Errorf("Invalid UUID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "invalid ID"})
	}

	if err := repository.DeleteOrganizationByID(id); err != nil {
		logger.Logger.Errorf("Failed to delete organization: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	logger.Logger.Infof("Organization deleted: %s", id)
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "message": "organization deleted"})
}
