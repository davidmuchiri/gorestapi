package users

import (
	"encoding/json"
	"net/http"

	"github.com/dinobambino7/gorestapi/utils"
)

//Register registers a user
func Register(res http.ResponseWriter, req *http.Request) {

	var user Account

	_ = json.NewDecoder(req.Body).Decode(&user)

	response := user.CreateAccount()

	utils.Response(res, response)

}

//Authenticate registers a user
func Authenticate(res http.ResponseWriter, req *http.Request) {

	var user Account
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		utils.Response(res, utils.Message(false, "invalid request", nil))
	}

	response := Login(user.UserEmail, user.UserPassword)

	utils.Response(res, response)

}
