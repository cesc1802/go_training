package models
import (
	_"fmt"

	_"github.com/jinzhu/gorm"
	"github.com/cesc1802/go_training/utils"
)

// User reflects users data from DB
type Users struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Max_todo  int  `json:"max___todo"`
	test     string `json:"test"`
}

func GetUsers() ([]*Users) {
	users := make([]*Users,0)
	err := GetDB().Table("Users").Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}
func (user *Users) Create() (map[string]interface{}) {

	//if resp, ok := contact.Validate(); !ok {
	//	return resp
	//}

	GetDB().Create(user)
	resp := utils.Message(true, "success")
	resp["user"] = user
	return resp
}