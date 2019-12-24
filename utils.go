package main

import (
	"os"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
 		port = "8080"
 		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
 	}
 	return ":" + port
 }

func InitRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir("./static"))))
	router.HandleFunc("/user", AddUser).Methods("POST")
	router.HandleFunc("/user/{pesel}", GetUser).Methods("GET")
	router.HandleFunc("/user/{pesel}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/book", AddBook).Methods("POST")
	router.HandleFunc("/book/{uniquecode}", GetBook).Methods("GET")
	router.HandleFunc("/books", GetBooks).Methods("GET")

	router.HandleFunc("/lend", Lend).Methods("POST")
	router.HandleFunc("/giveBack", Archivise).Methods("POST")
	
	return router
}

func PanicErr(w http.ResponseWriter, err error) {
	if err != nil {
		RespondWithJSON(w, http.StatusServiceUnavailable, nil)
		panic(err)
	}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Println(err)
		RespondWithJSON(w, http.StatusBadRequest, nil)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}