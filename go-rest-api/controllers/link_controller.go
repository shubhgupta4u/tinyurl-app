package controllers

import (
	logger "go-rest-api/middlewares"
	"go-rest-api/models"
	repository "go-rest-api/repositories"
	"go-rest-api/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Create
func CreateLinkHandler(c echo.Context) error {
	var link models.Link
	if err := c.Bind(&link); err != nil {
		logger.Logger.WithField("error", err.Error()).Error("Failed to bind request")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "Invalid request"})
	}

	if !link.CustomAlias || link.ShortCode == "" {
		link.ShortCode = utils.GenerateRandomShortCode(6)
		logger.Logger.WithField("short_code", link.ShortCode).Debug("Generated random short code")
	}

	if err := repository.CreateLink(&link); err != nil {
		logger.Logger.WithField("error", err.Error()).Error("Failed to create link")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	logger.Logger.WithFields(logrus.Fields{
		"id":         link.ID,
		"short_code": link.ShortCode,
		"user_id":    link.UserID,
		"org_id":     link.OrgID,
	}).Info("Link created successfully")

	return c.JSON(http.StatusCreated, map[string]interface{}{"status": "success", "data": link})
}

// Get
func GetLinkHandler(c echo.Context) error {
	shortCode := c.Param("short_code")
	link, err := repository.GetLinkByShortCode(shortCode)
	if err != nil {
		logger.Logger.WithField("error", err.Error()).Error("Failed to get link")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}
	if link == nil {
		logger.Logger.WithField("short_code", shortCode).Info("Link not found or expired")
		return c.JSON(http.StatusNotFound, map[string]interface{}{"status": "error", "message": "Link not found or expired"})
	}
	logger.Logger.WithField("short_code", shortCode).Info("Link retrieved successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "data": link})
}

// Delete
func DeleteLinkHandler(c echo.Context) error {
	shortCode := c.Param("short_code")
	if err := repository.DeleteLinkByShortCode(shortCode); err != nil {
		logger.Logger.WithField("error", err.Error()).Error("Failed to delete link")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}
	logger.Logger.WithField("short_code", shortCode).Info("Link deleted successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "message": "Deleted successfully"})
}

// Search
func SearchLinksHandler(c echo.Context) error {
	userIDStr := c.QueryParam("user_id")
	orgIDStr := c.QueryParam("org_id")
	tag := c.QueryParam("tag")

	var userID, orgID *uuid.UUID
	if userIDStr != "" {
		id, _ := uuid.Parse(userIDStr)
		userID = &id
	}
	if orgIDStr != "" {
		id, _ := uuid.Parse(orgIDStr)
		orgID = &id
	}

	links, err := repository.SearchLinks(userID, orgID, tag)
	if err != nil {
		logger.Logger.WithField("error", err.Error()).Error("Failed to search links")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "message": err.Error()})
	}

	logger.Logger.WithFields(logrus.Fields{
		"user_id": userID,
		"org_id":  orgID,
		"tag":     tag,
		"count":   len(links),
	}).Info("Search executed successfully")

	return c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "data": links})
}

// Update link
func UpdateLinkHandler(c echo.Context) error {
	id := c.Param("id")
	var link models.Link
	if err := c.Bind(&link); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}
	parsedID, _ := uuid.Parse(id)
	link.ID = uint64(parsedID.ID())
	if err := repository.UpdateLink(&link); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, link)
}

// Patch link
func PatchLinkHandler(c echo.Context) error {
	id := c.Param("id")
	updates := make(map[string]interface{})
	if err := c.Bind(&updates); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}
	parsedID, _ := uuid.Parse(id)
	if err := repository.PatchLink(uint64(parsedID.ID()), updates); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "Patched successfully")
}
