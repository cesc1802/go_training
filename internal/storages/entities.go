package storages

import (
	"time"
	"gorm.io/gorm"
)

// Task reflects tasks in DB
type Task struct {
	gorm.Model
	ID          string    `json:"id" gorm:"size:64;not null"`
	Content     string    `json:"content" gorm:"not null"`
	UserID      string    `json:"user_id" gorm:"size:64;not null"`
	CreatedDate time.Time `json:"created_date" gorm:"not null"`
}

// User reflects users data from DB
type User struct {
	gorm.Model
	ID       string `gorm:"size:64;not null"`
	Password string `gorm:"not null"`
	MaxTodo  int    `gorm:"default:5;not null"`
	Task     []Task `gorm:"foreignKey:UserID;not null"`
}

func Migration(db *gorm.DB) {
	// check table exist or not
	if !db.Migrator().HasTable(&User{}) && !db.Migrator().HasTable("users") {
		err := db.Migrator().CreateTable(&User{})
		if err != nil {
			panic("user table migration error")
		}
	}

	if !db.Migrator().HasTable(&Task{}) && !db.Migrator().HasTable("tasks") {
		err := db.Migrator().CreateTable(&Task{})
		if err != nil {
			panic("task table migration error")
		}
	}

}
