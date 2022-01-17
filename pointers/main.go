package main

import "log"

func main(){
	var str string
	str = "Green"

	log.Println("str is set to :", str)
	changeUsingPointer(&str)
	log.Println("After changing value with pointers :", str)
}

func changeUsingPointer(s *string){
	newValue := "Red"
	*s = newValue
}