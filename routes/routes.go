package routes

import (
	"todo/controllers"
	"todo/core/router"
	mw "todo/middlewares"
)

type Routes struct {
	Router *router.Router
}

var DefaultCtrl = controllers.DefaultController{}
var authCtrl = controllers.AuthController{}
var todoCtrl = controllers.TodoController{}

func (rts *Routes) InitializeRoutes() {
	/** Auth user routes */
	rts.Router.POST("/signup", authCtrl.Signup)
	rts.Router.POST("/signin", authCtrl.Signin)

	/** Todo routes */
	rts.Router.Middlewares(mw.AuthMiddleware).GET("/todos", todoCtrl.GetTodos)
	rts.Router.Middlewares(mw.AuthMiddleware).GET("/todos/{id:[0-9]+}", todoCtrl.DetailTodo)
	rts.Router.Middlewares(mw.AuthMiddleware).POST("/todos", todoCtrl.CreateTodo)
	rts.Router.Middlewares(mw.AuthMiddleware).PUT("/todos/{id:[0-9]+}", todoCtrl.UpdateTodo)
	rts.Router.Middlewares(mw.AuthMiddleware).DELETE("/todos/{id:[0-9]+}", todoCtrl.DeleteTodo)
	rts.Router.
		Middlewares(mw.AuthMiddleware).
		PATCH("/todos/{id:[0-9]+}/{status}", todoCtrl.UpdateStatus)

	/** Default Route Handler */
	rts.Router.Default(DefaultCtrl.PageNotFound)
}
