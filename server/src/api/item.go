package api

import (
	model_item "menagment-app-2/src/model/item"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *APIServer) CreateItem(c *gin.Context) {
	var req model_item.CreateItemRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.CreateItem(req, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item created successfully"})
}

func (a *APIServer) GetItems(c *gin.Context) {
	organization_id, exist := c.Get("organization_id")
	pageSizeInt, pageNumberInt, err := getPaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page_size or page_number"})
		return
	}

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	items, err := a.db.GetItems(organization_id.(string), pageSizeInt, pageNumberInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (a *APIServer) GetItem(c *gin.Context) {
	item_id := c.Param("item_id")
	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	item, err := a.db.GetItemById(item_id, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)

}

func (a *APIServer) UpdateItem(c *gin.Context) {
	var req model_item.UpdateItemRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item_id := c.Param("item_id")
	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.UpdateItem(item_id, organization_id.(string), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
}

func (a *APIServer) DeleteItem(c *gin.Context) {
	item_id := c.Param("item_id")
	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.DeleteItem(item_id, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func (a *APIServer) AssignItemToPackage(c *gin.Context) {
	var req model_item.AssignItemToPackageRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item_id := c.Param("item_id")
	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.AssignItemToPackage(item_id, organization_id.(string), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
}
