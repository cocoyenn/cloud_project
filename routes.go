package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
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
	
	RespondWithJSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User) RETURN user.name AS Name, user.surname AS Surname, user.age As Age, user.country as Country;" ,
		Result:    &users ,
	}

	err := Db.Cypher(&cq)
	PanicErr(w, err)
	
	RespondWithJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request){
	var user User
	var result []User

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return;
	}

	json.Unmarshal(reqBody, &user)

	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User) WHERE user.pesel = {pesel} RETURN user.name AS Name, user.surname AS Surname, user.age As Age, user.country as Country",
		Parameters: neoism.Props{"pesel": user.Pesel },
		Result:     &result,
	}
	for res in result:
		log.Println(res.name);
	err = Db.Cypher(&cq)
	PanicErr(w, err)

	RespondWithJSON(w, http.StatusOK, result)
}

func AddBook(w http.ResponseWriter, r *http.Request){
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

	RespondWithJSON(w, http.StatusCreated, book)
}

func GetBooks(w http.ResponseWriter, r *http.Request){
	var books []Book

	cq := neoism.CypherQuery{
		Statement: "MATCH (book:Book) RETURN book.title AS Title, book.type AS Type, book.uniquecode As UniqueCode ;" ,
		Result:    &books ,
	}

	err := Db.Cypher(&cq)
	PanicErr(w, err)
	
	RespondWithJSON(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	var book Book
	var books []Book

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return;
	}

	json.Unmarshal(reqBody, &book)

	cq := neoism.CypherQuery{
		Statement: "MATCH (book:Book) WHERE book.uniquecode = {uniquecode} RETURN book.title AS Title, book.type AS Type, book.uniquecode As UniqueCode",
		Parameters: neoism.Props{"uniquecode": book.UniqueCode },
		Result:     &books,
	}

	err = Db.Cypher(&cq)
	PanicErr(w, err)
	
	RespondWithJSON(w, http.StatusOK, books)
}


func DeleteUser(w http.ResponseWriter, r *http.Request){
	var user User
	var result []User

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return;
	}

	json.Unmarshal(reqBody, &user)

	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User) WHERE user.pesel = {pesel} DELETE user",
		Parameters: neoism.Props{"pesel": user.Pesel },
		Result:     &result,
	}

	err = Db.Cypher(&cq)
	PanicErr(w, err)
	
	RespondWithJSON(w, http.StatusNoContent, user)
}

func Lend(w http.ResponseWriter, r *http.Request){
	var lendHelper LendHelper
	res0 := []struct { N neoism.Node }{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return;
	}
	json.Unmarshal(reqBody, &lendHelper)
	
	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User {pesel: {pesel}}) MATCH (book:Book {uniquecode: {uniquecode}}) CREATE (user)-[rel:BORROWED]->(book)",
		Parameters: neoism.Props{"pesel": lendHelper.Pesel, "uniquecode" : lendHelper.UniqueCode},
		Result:     &res0,
	}

	err = Db.Cypher(&cq)
	PanicErr(w, err)
	
	RespondWithJSON(w, http.StatusCreated, nil)
}

func Archivise(w http.ResponseWriter, r *http.Request){
	var lendHelper LendHelper
	res0 := []struct { N neoism.Node }{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return;
	}
	json.Unmarshal(reqBody, &lendHelper)
	
	cq := neoism.CypherQuery{
		Statement: "MATCH (user:User {pesel: {pesel}}) MATCH (book:Book {uniquecode: {uniquecode}}) CREATE (user)-[rel:RETURNED]->(book)",
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