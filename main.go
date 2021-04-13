package main

import (
	"log"
	"net/http"
	"users-api/group"
	"users-api/user"
	"users-api/util"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", user.ListUsers)
	myRouter.HandleFunc("/user", user.SaveSingleUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", user.RemoveUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", user.FindSingleUser)
	myRouter.HandleFunc("/groups", group.ListGroups)
	myRouter.HandleFunc("/group", group.SaveSingleGroup).Methods("POST")
	myRouter.HandleFunc("/group/{id}", group.RemoveGroup).Methods("DELETE")
	myRouter.HandleFunc("/group/{id}", group.FindSingleGroup)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	util.OpenConnectionAndSetupSchema()
	handleRequests()
}