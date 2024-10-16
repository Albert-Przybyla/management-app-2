package api

import (
	model_storage "menagment-app-2/src/model/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *APIServer) CreateStorage(c *gin.Context) {
	var req model_storage.CreateStorageRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization_id, exist := c.Get("organization_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.CreateStorage(req, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Storage created successfully"})
}
