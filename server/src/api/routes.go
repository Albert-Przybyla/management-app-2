package api

func (a *APIServer) Routes() {
	a.engine.POST("/login", a.Login)
	a.engine.POST("/user", a.CreateUser)
}
