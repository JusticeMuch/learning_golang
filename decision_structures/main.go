package main

import "log"

func main() {
	//decision structures work same as other languages minus the brackets arounf the conditionals

	myVar := "cat"
	switch myVar {
	case "cat":
		log.Println("Is cat")

	case "dog":
		log.Println("Is dog")

	case "bear":
		log.Println("Is bear")
	}
}
