// @title Task & User Microservice API
// @version 1.0
// @description This is a microservice for managing tasks and users.
// @host localhost:8080
// @BasePath /

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/swaggo/http-swagger"
	_ "microservice/docs"

	"microservice/database"

	taskHandler "microservice/handler/task"
	userHandler "microservice/handler/user"

	taskService "microservice/service/task"
	userService "microservice/service/user"

	taskStore "microservice/store/task"
	userStore "microservice/store/user"
)

func main() {
	// Connect to the database
	db, err := database.Databasconnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return
	}

	fmt.Println(" Connected to the database successfully")

	tStore := taskStore.New(db)
	tService := taskService.New(tStore)
	tHandler := taskHandler.New(tService)

	http.HandleFunc("POST /task", tHandler.Addtask)
	http.HandleFunc("GET /task", tHandler.Getalltask)
	http.HandleFunc("GET /task/{id}", tHandler.Gettaskbyid)
	http.HandleFunc("PATCH /task/{id}", tHandler.Completetask)
	http.HandleFunc("DELETE /task/{id}", tHandler.Deletetask)

	uStore := userStore.New(db)
	uService := userService.New(uStore)
	uHandler := userHandler.New(uService)

	http.HandleFunc("POST /user", uHandler.AddUser)
	http.HandleFunc("GET /user", uHandler.GetAllUsers)
	http.HandleFunc("GET /user/{id}", uHandler.GetUserByID)
	http.HandleFunc("DELETE /user/{id}", uHandler.DeleteUserByID)
	http.HandleFunc("DELETE /user", uHandler.DeleteAllUsers)

	fmt.Println(" Swagger available at: http://localhost:8080/swagger/")
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/swagger.json"),
	))

	port := ":8080"
	fmt.Println(" Server is listening on port", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(errors.New(" Failed to start server: port may be in use"))
	}
}
