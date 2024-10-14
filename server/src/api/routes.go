package api

func (a *APIServer) Routes() {

	a.engine.POST("/login", a.Login)
	a.engine.POST("/user", a.CreateUser)
	a.engine.GET("/user", a.validateToken(), a.GetUser)
}
