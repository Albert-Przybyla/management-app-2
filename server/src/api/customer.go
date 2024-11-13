package api

import (
	model_customer "menagment-app-2/src/model/customer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *APIServer) CreateCustomer(c *gin.Context) {
	var req model_customer.CreateCustomerRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization_id, exist := c.Get("organization_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.CreateCustomer(req, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer created successfully"})
}

func (a *APIServer) GetCustomers(c *gin.Context) {
	organization_id, exist := c.Get("organization_id")
	pageSizeInt, pageNumberInt, err := getPaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page_size or page_number"})
		return
	}

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	items, err := a.db.GetCustomers(organization_id.(string), pageSizeInt, pageNumberInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (a *APIServer) GetCustomer(c *gin.Context) {
	customer_id := c.Param("customer_id")
	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	item, err := a.db.GetCustomer(customer_id, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)

}

func (a *APIServer) DeleteCustomer(c *gin.Context) {
	customer_id := c.Param("customer_id")
	organization_id, exist := c.Get("organization_id")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	err := a.db.DeleteCustomer(customer_id, organization_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
