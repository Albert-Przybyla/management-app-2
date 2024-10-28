package api

import (
	model_operations "menagment-app-2/src/model/operations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *APIServer) CreateReturn(c *gin.Context) {
	var req model_operations.CreateRetrunRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization_id, exist := c.Get("organization_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.CreateReturn(req, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
