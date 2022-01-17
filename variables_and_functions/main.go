package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var saySomething string
	var i int
	var say string

	saySomething = "Bye World"

	fmt.Println(saySomething)

	i = 7

	fmt.Println("i is = ", i)

	say, saySomething = saySome()

	fmt.Println("say something : ", say, saySomething)
}

func saySome() (string, string) {
	return "something", "buyer"
}
