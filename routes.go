package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmcvetta/neoism"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user User
	res0 := []struct { N neoism.Node }{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return;
	}
	json.Unmarshal(reqBody, &user)
	
	cq := neoism.CypherQuery{
		Statement: "CREATE (n:User {name: {name}, age: {age}, surname:{surname}, country:{country}, pesel: {pesel}}) RETURN n",
		Parameters: neoism.Props{"name": user.Name, "surname" : user.Surname, "age" : user.Age, "country" :user.Country, "pesel" :user.Pesel},
		Result:     &res0,
	}

	err = Db.Cypher(&cq)
	PanicErr(w, err)

	if(len(res0) == 0){
		RespondWithJSON(w, http.StatusForbidden, nil)
	}
	
	RespondWithJSON(w, http.StatusCreated, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var result []UniqueCodeHelper
	userPesel, _ := mux.Vars(r)["pesel"]

	res0 := []struct { N neoism.Node }{}
	cq_check := neoism.CypherQuery{
		Statement: "MATCH (user:User{pesel: {pesel}}) RETURN user",

		Parameters: neoism.Props{"pesel": userPesel },
		Result:     &res0,
	}
	err := Db.Cypher(&cq_check)
	
	if(len(res0) == 0) {
		RespondWithJSON(w, http.StatusNotFound, nil)
		return;
	}

	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User{pesel: {pesel}})-[r:BORROWED|:RETURNED]-> (book:Book) RETURN book.title AS Title, book.type AS Type, book.uniquecode As UniqueCode, type(r) as State",

		Parameters: neoism.Props{"pesel":userPesel },
		Result:     &result,
	}

	err = Db.Cypher(&cq)
	PanicErr(w, err)
	log.Println(len(result))
	RespondWithJSON(w, http.StatusOK, result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var result []User

	userPesel, _ := mux.Vars(r)["pesel"]

	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User{pesel: {pesel}}) OPTIONAL MATCH (user)-[r:RETURNED]-() DELETE user, r",
		Parameters: neoism.Props{"pesel": userPesel },
		Result:     &result,
	}

	err := Db.Cypher(&cq)
	if(err != nil){
		var msg Message 
		msg.MSG = err.Error()
		RespondWithJSON(w, http.StatusForbidden, msg)
		log.Println(err)
	} else {
		RespondWithJSON(w, http.StatusNoContent, user)
	}
	

}

func AddBook(w http.ResponseWriter, r *http.Request ){
	var book Book
	res0 := []struct { N neoism.Node }{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return;
	}
	json.Unmarshal(reqBody, &book)

	cq := neoism.CypherQuery{
		Statement: "CREATE (n:Book {title: {title}, type: {type}, uniquecode: {uniquecode}}) RETURN n",
		Parameters: neoism.Props{"title": book.Title, "type" : book.Type, "uniquecode": book.UniqueCode},
		Result:     &res0,
	}

	Db.Cypher(&cq)
	PanicErr(w, err)

	if(len(res0) == 0){
		RespondWithJSON(w, http.StatusForbidden, nil)
	}

	RespondWithJSON(w, http.StatusCreated, book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var books  []UserBookHealper
	UniqueCode, _ := mux.Vars(r)["uniquecode"]
	res0 := []struct { N neoism.Node }{}
	cq_check := neoism.CypherQuery{
		Statement: "MATCH (book:Book {uniquecode: {uniquecode}}) RETURN book",

		Parameters: neoism.Props{"uniquecode": UniqueCode },
		Result:     &res0,
	}
	err := Db.Cypher(&cq_check)
	
	if(len(res0) == 0) {
		RespondWithJSON(w, http.StatusNotFound, nil)
		return;
	}

	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User)-[r:BORROWED|:RETURNED]-> (book:Book {uniquecode: {uniquecode}}) RETURN user.name AS Name, user.surname AS Surname, user.pesel As Pesel, type(r) as State",
		Parameters: neoism.Props{"uniquecode": UniqueCode },
		Result:     &books,
	}
	
	err = Db.Cypher(&cq)
	PanicErr(w, err)
	RespondWithJSON(w, http.StatusOK, books)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books  []Book	
	cq_check := neoism.CypherQuery{
		Statement: "MATCH (book:Book) RETURN  book.title AS Title, book.type AS Type, book.uniquecode As UniqueCode",
		Result:     &books,
	}
	err := Db.Cypher(&cq_check)
	PanicErr(w, err)
	RespondWithJSON(w, http.StatusOK, books)
}

func Lend(w http.ResponseWriter, r *http.Request) {
	var lendHelper LendHelper
	res0 := []struct { N neoism.Node }{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return;
	}
	json.Unmarshal(reqBody, &lendHelper)

	cq_check := neoism.CypherQuery{
		Statement: "MATCH (user:User{pesel: {pesel}}) RETURN user",

		Parameters: neoism.Props{"pesel": lendHelper.Pesel },
		Result:     &res0,
	}

	err = Db.Cypher(&cq_check)
	if(len(res0) == 0) {
		RespondWithJSON(w, http.StatusNotFound, nil)
		return;
	}

	cq_check = neoism.CypherQuery{
		Statement: "MATCH (book:Book {uniquecode: {uniquecode}}) RETURN book",

		Parameters: neoism.Props{"uniquecode": lendHelper.UniqueCode },
		Result:     &res0,
	}

	err = Db.Cypher(&cq_check)
	
	if(len(res0) == 0) {
		RespondWithJSON(w, http.StatusNotFound, nil)
		return;
	}
	
	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User {pesel: {pesel}}) MATCH (book:Book {uniquecode: {uniquecode}}) CREATE (user)-[r:BORROWED]->(book) RETURN type(r)",
		Parameters: neoism.Props{"pesel": lendHelper.Pesel, "uniquecode" : lendHelper.UniqueCode},
		Result:     &res0,
	}
	
	err = Db.Cypher(&cq)
	PanicErr(w, err)
	
	RespondWithJSON(w, http.StatusCreated, nil)
}

func Archivise(w http.ResponseWriter, r *http.Request) {
	var lendHelper LendHelper
	res0 := []struct { N neoism.Node }{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return;
	}
	json.Unmarshal(reqBody, &lendHelper)

	cq_check := neoism.CypherQuery{
		Statement: "MATCH (user:User{pesel: {pesel}}) RETURN user",

		Parameters: neoism.Props{"pesel": lendHelper.Pesel },
		Result:     &res0,
	}

	err = Db.Cypher(&cq_check)
	if(len(res0) == 0) {
		RespondWithJSON(w, http.StatusNotFound, nil)
		return;
	}

	cq_check = neoism.CypherQuery{
		Statement: "MATCH (book:Book {uniquecode: {uniquecode}}) RETURN book",

		Parameters: neoism.Props{"uniquecode": lendHelper.UniqueCode },
		Result:     &res0,
	}
	
	err = Db.Cypher(&cq_check)
	
	if(len(res0) == 0) {
		RespondWithJSON(w, http.StatusNotFound, nil)
		return;
	}
	
	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User {pesel: {pesel}}) MATCH (book:Book {uniquecode: {uniquecode}}) CREATE (user)-[rel:RETURNED]->(book) RETURN user",
		Parameters: neoism.Props{"pesel": lendHelper.Pesel, "uniquecode" : lendHelper.UniqueCode},
		Result:     &res0,
	}

	err = Db.Cypher(&cq)
	PanicErr(w, err)

	cq_delete := neoism.CypherQuery{
		Statement: "MATCH (user:User {pesel: {pesel}})-[rel:BORROWED]-> (book:Book {uniquecode: {uniquecode}}) DELETE rel",
		Parameters: neoism.Props{"pesel": lendHelper.Pesel, "uniquecode" : lendHelper.UniqueCode},
		Result:     &res0,
	}

	err = Db.Cypher(&cq_delete)
	PanicErr(w, err)
	
	RespondWithJSON(w, http.StatusCreated, nil)
}