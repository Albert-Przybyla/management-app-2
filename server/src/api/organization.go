package api

import (
	model_organization "menagment-app-2/src/model/organization"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *APIServer) CreateOrganization(c *gin.Context) {
	var req model_organization.CreateOrganizationRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.db.CreateOrganization(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (a *APIServer) GetOrganization(c *gin.Context) {
	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get organization ID"})
	}

	res, err := a.db.GetOrganizationById(organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
