package main

import (
	"log"
)

// var s = "six"

// type User struct {
// 	FirstName 	string
// 	LastName 	string
// 	PhoneNumber string
// 	Age			int
// 	BirthDate 	time.Time
// }

// func main(){
// 	user := User{
// 		FirstName: "Trevor",
// 		LastName: "Mornes",
// 		PhoneNumber: "454545454",
// 	}

// 	log.Println(user.FirstName, user.LastName, user.BirthDate)
// }

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

type myStruct struct {
	FirstName string
}

func main(){
	var myVar myStruct
	myVar.FirstName = "jfdsf";

	myVar2 := myStruct{
		FirstName: "kkok",
	}

	log.Println("myVar is set to ", myVar.printFirstName())
	log.Println("myVar2 is set to ", myVar2.printFirstName())
}
