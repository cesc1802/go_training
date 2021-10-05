package storages

import (
	"gorm.io/gorm"
)

type TaskRepo interface {
	FindById(id string)
	Save(t *Task)
	Update(t *Task)
}

type taskRepo struct {
}

func NewTaskRepo() *taskRepo {
	return &taskRepo{}
}

func (tr *taskRepo) FindById(db *gorm.DB, id string, t *Task) {
	db.Raw("SELECT * FROM tasks WHERE id = ?", id).Scan(&t)
}

func (tr *taskRepo) Save(db *gorm.DB, t *Task) *Task {
	db.Create(t)
	return t
}

func (tr *taskRepo) UpdateContent(db *gorm.DB, t *Task) *Task {
	db.Save(&t)
	return t
}
