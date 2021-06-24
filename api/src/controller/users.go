package controller

import "net/http"

//Create user
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("CreateUser"))
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
