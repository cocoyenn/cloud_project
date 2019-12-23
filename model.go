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

type UserBookHealper struct {
		Name       	string `json:"Name"`
		Surname 	string `json:"Surname"`
		Pesel		string `json:"Pesel"`
		State	 	string `json:"State"`
	
}

type UniqueCodeHelper struct {
	Title		string `json:"Title"`
	Type		string `json:"Type"`
	UniqueCode 	string `json:"UniqueCode"`
	State	 	string `json:"State"`
}

type Message struct {
	MSG		string `json:"Msg"`
}
