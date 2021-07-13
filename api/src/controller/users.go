package controller

import (
	"api/src/answers"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//Create user
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Err(rw, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		answers.Err(rw, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {
		answers.Err(rw, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Err(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		answers.Err(rw, http.StatusInternalServerError, err)
		return
	}
	user.ID = userID
	answers.JSON(rw, http.StatusCreated, user)
}

//Find users
func FindUsers(rw http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		answers.Err(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	users, err := repository.Find(nameOrNick)

	if err != nil {
		answers.Err(rw, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(rw, http.StatusOK, users)

}

//Find user by id
func FindUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		answers.Err(rw, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Err(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	user, err := repository.FindById(userID)

	if err != nil {
		answers.Err(rw, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(rw, http.StatusOK, user)
}

//Update user
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("UpdateUser"))
}

//Delete user
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("DeleteUser"))
}
