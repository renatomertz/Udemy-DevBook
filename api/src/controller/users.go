package controller

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Create user
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewUsersRepository(db)
	repository.Create(user)
}

//Find users
func FindUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("FindUsers"))
}

//Find user by id
func FindUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("FindUser"))
}

//Update user
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("UpdateUser"))
}

//Delete user
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("DeleteUser"))
}
