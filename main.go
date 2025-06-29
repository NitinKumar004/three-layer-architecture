package main

import (
	"errors"
	"fmt"
	"log"
	"microservice/database"
	handlertask "microservice/handler/task"
	handleruser "microservice/handler/user"
	servicetask "microservice/service/task"
	serviceuser "microservice/service/user"
	storetask "microservice/store/task"
	storeuser "microservice/store/user"
	"net/http"
)

func main() {

	db, err := database.Databasconnection()
	if err != nil {

		log.Fatal(err)
		return
	}
	//manage task function here
	store := storetask.New(db)
	sv := servicetask.New(store)
	h := handlertask.New(sv)

	http.HandleFunc("POST /task", h.Addtask)
	http.HandleFunc("GET /task", h.Getalltask)
	http.HandleFunc("GET /task/{id}", h.Gettaskbyid)
	http.HandleFunc("PATCH /task/{id}", h.Completetask)
	http.HandleFunc("DELETE /task/{id}", h.Deletetask)

	//manage user function here
	storeuser := storeuser.New(db)
	svuser := serviceuser.New(storeuser)
	huser := handleruser.New(svuser)
	http.HandleFunc("POST /user", huser.AddUser)
	http.HandleFunc("GET /user", huser.GetAllUsers)
	http.HandleFunc("GET /user/{id}", huser.GetUserByID)
	http.HandleFunc("DELETE /user/{id}", huser.DeleteUserByID)
	http.HandleFunc("DELETE /user", huser.DeleteAllUsers)

	fmt.Println("connected to database successfully")
	fmt.Println("listening on port number : ", 8080)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		err = errors.New("error to openning the port may be someoneuse using the port")
		if err != nil {
			return
		}
		return
	}

}
