package routes

import (
	"github.com/Tijanieneye10/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {

	route.GET("/todos", controllers.GetTodos)
	route.POST("/todos", controllers.CreateTodo)
	route.DELETE("/todos/:id", controllers.DeleteTodo)
	route.PUT("/todos/completed/:id", controllers.MarkAsCompletedTodo)
	route.PUT("/todos/uncompleted/:id", controllers.MarkAsUnCompletedTodo)
}
