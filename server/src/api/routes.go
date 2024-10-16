package api

func (a *APIServer) Routes() {

	// Login and register routes
	a.engine.POST("/login", a.Login)
	a.engine.POST("/register", a.Register)

	// User routes
	a.engine.POST("/user", a.validateToken(), a.CreateUser)
	a.engine.GET("/user", a.validateToken(), a.GetUser)

	// Organization routes
	a.engine.POST("/organization", a.CreateOrganization)

	// Storage routes
	a.engine.POST("/storage", a.validateToken(), a.CreateStorage)
}
