package api

func (a *APIServer) Routes() {

	// Login and register routes
	a.engine.POST("/login", a.Login)
	a.engine.POST("/register", a.Register)

	// User routes
	a.engine.POST("/user", a.validateToken(), a.CreateUser)
	a.engine.GET("/user", a.validateToken(), a.GetUser)
	a.engine.PUT("/user", a.validateToken(), a.UpdateUser)
	a.engine.PATCH("/user", a.validateToken(), a.ChangePassword)

	// Organization routes
	a.engine.POST("/organization", a.CreateOrganization)
	a.engine.GET("/organization", a.validateToken(), a.GetOrganization)

	// Storage routes
	a.engine.POST("/storage", a.validateToken(), a.CreateStorage)

	// Item routes
	a.engine.POST("/item", a.validateToken(), a.CreateItem)
	a.engine.GET("/item", a.validateToken(), a.GetItems)
	a.engine.GET("/item/:item_id", a.validateToken(), a.GetItem)
	a.engine.PUT("/item/:item_id", a.validateToken(), a.UpdateItem)
	a.engine.DELETE("/item/:item_id", a.validateToken(), a.UpdateItem)
	a.engine.PATCH("/item/:item_id", a.validateToken(), a.AssignItemToPackage)
}
