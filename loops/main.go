package main

import "log"

func main(){
	// for i := 0; i < 5; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"dog", "cat", "fish", "cow"}

	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	type User struct {
		FirstName string
		LastName string
	}

	var users []User

	users = append(users, User{"Jim", "Jones"})
	users = append(users, User{"James", "Jones"})
	users = append(users, User{"Sam", "Smith"})

	for index, user := range users{
		log.Println(index, user)
	}
}