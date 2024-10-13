package api

import (
	model_user "menagment-app-2/src/model/user"

	"github.com/gin-gonic/gin"
)

func (a *APIServer) Login(c *gin.Context) {
	var req model_user.LoginUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// err := a.db.Login(req.Username, req.Password)
	// if err != nil {
	// 	c.JSON(401, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(200, model_user.TokenUserResponse{Token: "jeeee"})
}
