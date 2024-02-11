package controllers

import (
	"github.com/Tijanieneye10/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTodos(context *gin.Context) {
	todos, err := models.GetTodos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Todos fetched successfully", "todos": todos})
}
func CreateTodo(context *gin.Context) {

	var todo *models.Todo
	err := context.ShouldBindJSON(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdTodo, err := todo.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Yeah is working fine", "data": createdTodo})
}

func MarkAsCompletedTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = models.MarkTodoCompleted(todoId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo marked as completed"})
}

func MarkAsUnCompletedTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = models.UndoTodo(todoId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo marked as uncompleted"})
}

func DeleteTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = models.DeleteTodo(todoId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
