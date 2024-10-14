package api

import (
	model_user "menagment-app-2/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Description Creates a new user account with provided details.
// @Tags user
// @Accept json
// @Produce json
// @Param createUserRequest body model_user.CreateUserRequest true "User data"
// @Success 200 {object} model_user.TokenUserResponse
// @Failure 400 {object} map[string]string
// @Router /create-user [post]

func (a *APIServer) CreateUser(c *gin.Context) {
	var req model_user.CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.db.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model_user.TokenUserResponse{Token: "sample-token"})
}

func (a *APIServer) GetUser(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}
	user, err := a.db.GetUserById(id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
