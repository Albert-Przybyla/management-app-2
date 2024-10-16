package api

func (a *APIServer) Routes() {

	a.engine.POST("/login", a.Login)
	a.engine.POST("/user", a.CreateUser)
	a.engine.GET("/user", a.validateToken(), a.GetUser)

	a.engine.POST("/organization", a.CreateOrganization)

	a.engine.POST("/storage", a.validateToken(), a.CreateStorage)
}
