package controllers

import (
	"net/http"
	"strings"

	dto_request "github.com/cesc1802/go_training/internal/dto/requests"
	"github.com/cesc1802/go_training/internal/services"
	"github.com/cesc1802/go_training/internal/storages"
	"github.com/gin-gonic/gin"
)

type taskController struct{}

func (tc taskController) Routes(r *gin.Engine) {
	r.POST("/tasks", createTask)
}

func createTask(c *gin.Context) {
	var task dto_request.Task
	c.BindJSON(&task)
	user := c.MustGet("USER").(storages.User)

	var contentInDay int
	var sb strings.Builder
	sb.WriteString("SELECT COUNT(*) FROM tasks ")
	sb.WriteString("JOIN users ON users.id = tasks.user_id ")
	sb.WriteString("WHERE CAST(tasks.created_at AS DATE) = CURRENT_DATE()")
	storages.Get().Raw(sb.String()).Scan(&contentInDay)
	if contentInDay > user.MaxTodo {
		c.String(http.StatusBadRequest, "Max todo in day")
		c.Abort()
		return
	}
	result := services.CreateTask(&task, user)
	c.JSON(http.StatusCreated, result)
}

func TaskController() taskController {
	return taskController{}
}
