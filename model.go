package main

type User struct {
	Name       	string `json:"Name"`
	Surname 	string `json:"Surname"`
	Age			string `json:"Age"`
	Country 	string `json:"Country"`
	Pesel		string `json:"Pesel"`
}


type Book struct {
	Title		string `json:"Title"`
	Type		string `json:"Type"`
	UniqueCode 	string `json:"UniqueCode"`
}

type LendHelper struct {
	Pesel		string `json:"Pesel"`
	UniqueCode 	string `json:"UniqueCode"`
}


// CREATE CONSTRAINT ON (n:Person)
// ASSERT n.name IS UNIQUE
