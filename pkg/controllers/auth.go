package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/9996rojit/backend_go/pkg/models"
	"github.com/9996rojit/backend_go/pkg/utils"
)

var UserData models.User

//
// func HandleHome(w http.ResponseWriter, r *http.Request) {
//
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(userDetail)
// }

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	registerUser := &models.User{}
	utils.BodyParser(r, registerUser)
	u := registerUser.CreateUser()
	jsonData, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "pkglication/json")
		w.Write([]byte("Error While converting json"))
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(jsonData)
}

//
// func HandleLogin(w http.ResponseWriter, r *http.Request) {
// 		var user login
// 		utils.BodyParser(r, &user)
//
//
//
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Header().Set("Content-Type", "pkglication/json")
// 		w.Write([]byte("Program is unable to parse req.body"))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	encodedUser, _ := json.Marshal(loginInfo)
// 	w.Write([]byte(encodedUser))
// }

// func HandleForgetPassword(w http.ResponseWriter, r *http.Request) {
//
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte("This is response that function forget password should respond to client"))
// }

// func HandleChangePassword(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.Write([]byte("This is response that server handle change password function should respond to client"))
// }
