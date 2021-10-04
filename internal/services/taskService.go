package services

import (
	"time"

	dto_request "github.com/cesc1802/go_training/internal/dto/requests"
	"github.com/cesc1802/go_training/internal/storages"
	"github.com/google/uuid"
)

func CreateTask(t *dto_request.Task, user storages.User) storages.Task {
	id := uuid.New()
	task := storages.Task{
		ID: id.String(),
		Content: t.Content,
		UserID: user.ID,
		CreatedDate: time.Now(),
	}
	storages.NewTaskRepo().Save(&task)
	return task
}