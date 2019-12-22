package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jmcvetta/neoism"
)

func InitRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir("./static"))))
	router.HandleFunc("/user", AddUser).Methods("POST")
	router.HandleFunc("/user", GetUser).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/delete", DeleteUser).Methods("DELETE")

	router.HandleFunc("/book", AddBook).Methods("POST")
	router.HandleFunc("/book", GetBook).Methods("GET")
	router.HandleFunc("/books", GetBooks).Methods("GET")

	router.HandleFunc("/lend", Lend).Methods("POST")
	router.HandleFunc("/giveBack", Archivise).Methods("POST")
	
	return router
}

var Db *neoism.Database

func main(){
	var err error = nil

	//var databasePath string = "http://neo4j:caro@localhost:7474"
	var databasePath string = "https://caro:b.HIBN6GItb7fD.fn3PKjJSXseIjyzl@hobby-oghlklmkakojgbkepalfbfdl.dbs.graphenedb.com:24780/db/data/"
	
	log.Print("Begining of initialization of server...\n")

	Db, err = neoism.Connect(databasePath) 
	if err != nil {
		log.Println("Connection to database failed. Exiting the program:\n", err)
		panic(err)
	}

	router := InitRouter()
	log.Println("Initialization end with succes: Server ready.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
