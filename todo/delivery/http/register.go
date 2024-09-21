package http

import (
	"github.com/gin-gonic/gin"
	"todo/todo"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc todo.UseCase) {
	h := NewHandler(uc)

	todos := router.Group("/todos")
	{
		todos.POST("", h.Create)
		todos.DELETE("", h.Delete)
		todos.GET("", h.Get)
		todos.PUT("", h.ChangeStatus)
	}
}
