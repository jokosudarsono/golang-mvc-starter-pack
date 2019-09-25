package controllers

import (
	"fmt"
	"net/http"
	"todo/core/response"
	"todo/models"
	"todo/utils/encryption"
	"todo/utils/jwt"
)

type AuthController struct{}

var authModel = models.AuthModel{}

func (ctrl *AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if firstname == "" {
		response.Err(w, http.StatusBadRequest, "Firstname cannot be empty!")
		return
	}

	if email == "" {
		response.Err(w, http.StatusBadRequest, "Email cannot be empty!")
		return
	}

	if password == "" {
		response.Err(w, http.StatusBadRequest, "Password cannot be empty!")
		return
	}

	user, _ := authModel.GetUserByEmail(email)
	if user != nil {
		response.Err(
			w,
			http.StatusBadRequest,
			"Email have been taken by anoter user",
			"user_exists",
		)

		return
	}

	hashPassword, _ := encryption.HashPassword(password)
	data, err := authModel.Signup(firstname, lastname, email, hashPassword)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err.Error())
		return
	}

	results := map[string]interface{}{"id": data}

	response.JSON(w, http.StatusOK, results)
}

func (ctrl *AuthController) Signin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" {
		response.Err(w, http.StatusBadRequest, "Email cannot be empty!")
		return
	}

	if password == "" {
		response.Err(w, http.StatusBadRequest, "Password cannot be empty!")
		return
	}

	user, err := authModel.GetUserByEmail(email)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err.Error())
		return
	}

	matched := encryption.CheckPasswordHash(password, fmt.Sprintf("%v", user["password"]))
	if !matched {
		response.Err(w, http.StatusUnauthorized, "Invalid credentials", "unauthorized")
		return
	}

	//	generate jwt token
	token, err := jwt.GenerateToken(user["id"].(int64), 100)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err.Error())
		return
	}

	results := map[string]interface{}{
		"token":      token,
		"id":         user["id"],
		"firstname":  user["firstname"],
		"lastname":   user["lastname"],
		"email":      user["email"],
		"created_at": user["created_at"],
		"updated_at": user["updated_at"],
	}

	response.JSON(w, http.StatusOK, results)
}
