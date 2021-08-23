package controllers

import (
	"encoding/json"
	_"encoding/json"

	"github.com/cesc1802/go_training/models"
	"github.com/cesc1802/go_training/utils"
	"net/http"
)

var GetUsers = func(w http.ResponseWriter, r *http.Request) {

	//id := r.Context().Value("user").(uint)
	data := models.GetUsers()
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	//user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	user := &models.Users{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	//contact.UserId = user
	resp := user.Create()
	utils.Respond(w, resp)
}