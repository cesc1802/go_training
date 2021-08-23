package controllers

import (
	"encoding/json"
	_"encoding/json"

	"github.com/cesc1802/go_training/models"
	"github.com/cesc1802/go_training/utils"
	"net/http"
)

var GetTask = func(w http.ResponseWriter, r *http.Request) {

	//id := r.Context().Value("user").(uint)
	data := models.GetTasks()
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var CreateTask = func(w http.ResponseWriter, r *http.Request) {

	//user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	task := &models.Task{}

	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	//contact.UserId = user
	resp := task.Create()
	utils.Respond(w, resp)
}