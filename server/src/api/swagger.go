package api

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *APIServer) Swagger() {
	// @title           API Example
	// @version         1.0
	// @description     To jest przykładowy opis API w Go z użyciem Swaggera.
	// @termsOfService  http://swagger.io/terms/
	// @contact.name    API Support
	// @contact.url     http://www.swagger.io/support
	// @contact.email   support@swagger.io
	// @license.name    Apache 2.0
	// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
	// @host            localhost:8080
	// @BasePath        /api/v1

	a.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
