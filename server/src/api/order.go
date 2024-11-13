package api

import (
	model_order "menagment-app-2/src/model/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *APIServer) CreateOrder(c *gin.Context) {
	var req model_order.CreateOrderRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization_id, exist := c.Get("organization_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.CreateOrder(req, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer created successfully"})
}

func (a *APIServer) GetOrders(c *gin.Context) {
	organization_id, exist := c.Get("organization_id")
	pageSizeInt, pageNumberInt, err := getPaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page_size or page_number"})
		return
	}

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	items, err := a.db.GetOrders(organization_id.(string), pageSizeInt, pageNumberInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (a *APIServer) GetOrder(c *gin.Context) {
	order_id := c.Param("order_id")
	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	item, err := a.db.GetOrder(order_id, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)

}

func (a *APIServer) DeleteOrder(c *gin.Context) {
	order_id := c.Param("order_id")
	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.DeleteOrder(order_id, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
