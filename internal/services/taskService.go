package services

import (
	"errors"
	"net/http"
	"strings"
	"time"

	dto_request "github.com/cesc1802/go_training/internal/dto/requests"
	"github.com/cesc1802/go_training/internal/storages"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTask(c *gin.Context, t *dto_request.Task, user storages.User) (storages.Task, error) {
	id := uuid.New()
	task := storages.Task{
		ID:          id.String(),
		Content:     t.Content,
		UserID:      user.ID,
		CreatedDate: time.Now(),
	}

	db := storages.Get()
	var contentInDay int

	var sb strings.Builder
	sb.WriteString("SELECT COUNT(*) FROM tasks ")
	sb.WriteString("JOIN users ON users.id = tasks.user_id ")
	sb.WriteString("WHERE CAST(tasks.created_at AS DATE) = CURRENT_DATE()")

	tx := db.Begin()
	storages.NewTaskRepo().Save(tx, &task)
	tx.Raw(sb.String()).Scan(&contentInDay)
	if contentInDay > user.MaxTodo {
		tx.Rollback()
		c.String(http.StatusBadRequest, "Max todo in day")
		c.Abort()
		return storages.Task{}, errors.New("max todo in day")
		} else {
		tx.Commit()
	}

	return task, nil
}

func UpdateContent(c *gin.Context) *storages.Task {
	db := storages.Get()
	id := c.Param("id")
	taskRepo := storages.NewTaskRepo()
	var task storages.Task
	taskRepo.FindById(db, id, &task)
	if task == (storages.Task{}) {
		return nil
	}
	updateContent := &dto_request.UpdateContent{}
	c.BindJSON(updateContent)
	task.Content = updateContent.Content
	return taskRepo.UpdateContent(db, &task)
}