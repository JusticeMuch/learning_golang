package main

import (
	"log"
	"testmod/helpers"
)

func main(){
	var test helpers.SomeType

	test.Test1 = "hello"
	test.Test2 = "Bye"

	log.Println(test)
}