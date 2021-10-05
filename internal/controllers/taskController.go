package controllers

import (
	"net/http"

	dto_request "github.com/cesc1802/go_training/internal/dto/requests"
	"github.com/cesc1802/go_training/internal/services"
	"github.com/cesc1802/go_training/internal/storages"
	"github.com/gin-gonic/gin"
)

type taskController struct{}

func (tc taskController) Routes(r *gin.Engine) {
	r.POST("/tasks", createTask)
	r.PATCH("/tasks/:id", updateTask)
}

func createTask(c *gin.Context) {
	var task dto_request.Task
	c.BindJSON(&task)
	user := c.MustGet("USER").(storages.User)
	result, err := services.CreateTask(c, &task, user)
	if err != nil {
		c.String(http.StatusBadRequest, "")
	} else {
		c.JSON(http.StatusCreated, result)
	}
}

func updateTask(c *gin.Context) {
	task := services.UpdateContent(c)	
	if task != (&storages.Task{}) {
		c.String(http.StatusBadRequest, "")
		c.Abort()
	}
	c.JSON(http.StatusAccepted, task)
}



func TaskController() taskController {
	return taskController{}
}
