package main

import "log"

type User struct {
	FirstName string
	LastName string
}

func main(){
	myMap := make(map[string]string)//basically keyed arrays 

	myMap["dog"] = "Samsom"
	myMap["other-dog"] = "Fido"

	myMap2 := make(map[string]User)

	me := User{
		FirstName: "lklk",
		LastName: "sdfdf",
	}

	myMap2["me"] = me

	var mySlice []string

	mySlice = append(mySlice, "Jim")//basically array
	mySlice = append(mySlice, "kkdfd")

	log.Println(mySlice)
	
	
}