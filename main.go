package main

import (
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jmcvetta/neoism"
)

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
 		port = "8080"
 		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
 	}
 	return ":" + port
 }

func InitRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", addCookie)
	router.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir("./static"))))
	router.HandleFunc("/user", AddUser).Methods("POST")
	router.HandleFunc("/user/{pesel}", GetUser).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/user/{pesel}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/book", AddBook).Methods("POST")
	router.HandleFunc("/book/{uniquecode}", GetBook).Methods("GET")
	router.HandleFunc("/books", GetBooks).Methods("GET")

	router.HandleFunc("/lend", Lend).Methods("POST")
	router.HandleFunc("/giveBack", Archivise).Methods("POST")
	
	return router
}

func addCookie(w http.ResponseWriter, r *http.Request) {
    cookie := http.Cookie{
        Name:    "SameSite",
        Value:  "none",
	}
	cookie2 := http.Cookie{
        Name:    "SameSite",
        Value:  "Secure",
    }
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}


var Db *neoism.Database

func main(){
	var err error = nil

	//var databasePath string = "http://neo4j:caro@localhost:7474"
	var databasePath string = "https://User1:b.ltlTUqioCmH1.zpLdQk4fhizAfsAv@hobby-oghlklmkakojgbkepalfbfdl.dbs.graphenedb.com:24780/db/data/"
	
	log.Print("Begining of initialization of server...\n")

	Db, err = neoism.Connect(databasePath) 
	if err != nil {
		log.Println("Connection to database failed. Exiting the program:\n", err)
		panic(err)
	}

	router := InitRouter()
	port := GetPort()
	log.Println("Initialization end with succes: Server ready." + port)
	log.Fatal(http.ListenAndServe(port, router))
}
