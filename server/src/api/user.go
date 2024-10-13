package api

import (
	model_user "menagment-app-2/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *APIServer) CreateUser(c *gin.Context) {
	var req model_user.CreateUserRequest

	// Walidacja danych wejściowych
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Wywołanie funkcji CreateUser z pliku database
	err := a.db.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Zwróć odpowiedź (np. token)
	c.JSON(http.StatusOK, model_user.TokenUserResponse{Token: "sample-token"})
}
