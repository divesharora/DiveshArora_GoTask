package main

import (
	"log"
	"DiveshArora_GoTask/db"
	"DiveshArora_GoTask/controllers"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql" 
	"github.com/gorilla/mux"

)


func main() {
		initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8000", router))
	

}
//router
func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.AddUsers).Methods("POST")
	router.HandleFunc("/get/nearme", controllers.GetUsersWithinDistance).Methods("GET")
	router.HandleFunc("/get/name", controllers.GetPersonByName).Methods("GET")
	router.HandleFunc("/createLikes", controllers.AddLikes).Methods("POST")
	router.HandleFunc("/get/matches", controllers.GetMatches).Methods("GET")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "secret",
			DB:         "stumble",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	
}