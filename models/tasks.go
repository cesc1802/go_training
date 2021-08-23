package models

import (
	_ "fmt"

	"github.com/cesc1802/go_training/utils"
	_ "github.com/jinzhu/gorm"
)
// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}


func GetTasks() ([]*Task) {
	task := make([]*Task,0)
	err := GetDB().Table("tasks").Find(&task).Error
	if err != nil {
		return nil
	}
	return task
}
func (task *Task) Create() (map[string]interface{}) {

	//if resp, ok := contact.Validate(); !ok {
	//	return resp
	//}

	GetDB().Create(task)

	resp := utils.Message(true, "success")
	resp["task"] = task
	return resp
}