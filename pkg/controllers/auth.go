package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/9996rojit/backend_go/pkg/models"
	"github.com/9996rojit/backend_go/pkg/utils"
)

var UserData models.User

// func HandleHome(w http.ResponseWriter, r *http.Request) {
//
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(userDetail)
//	}
type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	registerUser := &models.User{}
	// utils.BodyParser(r, registerUser)
	err := json.NewDecoder(r.Body).Decode(registerUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	u := registerUser.CreateUser()
	jsonData, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "pkglication/json")
		w.Write([]byte("Error While converting json"))
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(jsonData)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var user Login
	json.NewDecoder(r.Body).Decode(&user)
	var password = user.Password
	newUser := models.FindUserByEmail(user.Email)
	if newUser != nil {
		if err := utils.ComparePassword(password, newUser.Password); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		token, err := utils.GenerateToken(newUser.Email, newUser.RoleId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to generate token"))
			return
		}
		var res Token
		res.TokenString = token
		res.Email = newUser.Email
		res.Role = newUser.RoleId

		response, _ := json.Marshal(res)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "pkglication/json")
		w.Write(response)
		return

	}
}

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
