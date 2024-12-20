package api

import (
	model_organization "menagment-app-2/src/model/organization"
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

func (a *APIServer) Register(c *gin.Context) {
	var req model_user.CreateOwnerRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.db.CreateOrganization(model_organization.CreateOrganizationRequest{
		Name: req.OrganizationName,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = a.db.CreateUser(model_user.CreateUser{
		Email:          req.Email,
		Password:       req.Password,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		OrganizationID: res.Id,
		Role:           "owner",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model_user.TokenUserResponse{Token: "sample-token"})
}

func (a *APIServer) CreateUser(c *gin.Context) {
	var req model_user.CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization_id, exist := c.Get("organization_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get organization ID"})
	}

	res, err := a.db.CreateUser(model_user.CreateUser{
		Email:          req.Email,
		Password:       req.Password,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		OrganizationID: organization_id.(string),
		Role:           req.Role,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (a *APIServer) GetUsers(c *gin.Context) {
	organization_id, exist := c.Get("organization_id")
	pageSizeInt, pageNumberInt, err := getPaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page_size or page_number"})
		return
	}

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
	}

	items, err := a.db.GetUsers(organization_id.(string), pageSizeInt, pageNumberInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (a *APIServer) GetUser(c *gin.Context) {
	user_id, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized"})
	}
	user, err := a.db.GetUserById(user_id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (a *APIServer) UpdateUser(c *gin.Context) {
	var req model_user.UpdateUserRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized"})
	}
	err := a.db.UpdateUser(user_id.(string), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (a *APIServer) ChangePassword(c *gin.Context) {
	// TODO: Check Password
	var req model_user.ChangePasswordRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized"})
	}
	err := a.db.UpdateUserPassword(user_id.(string), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
