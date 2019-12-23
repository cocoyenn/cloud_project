package main

import (
	"log"
	"net/http"
	"github.com/jmcvetta/neoism"
)

func LogError(err error) {
	if(err != nil) {
		log.Println(err.Error())
	} 
}

var Db *neoism.Database

func main(){
	var err error = nil
	var databasePath string = "https://caro:b.HIBN6GItb7fD.fn3PKjJSXseIjyzl@hobby-oghlklmkakojgbkepalfbfdl.dbs.graphenedb.com:24780/db/data/"
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
