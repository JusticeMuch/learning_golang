package main

import "fmt"

type Animal interface{
	Says() string
	NumberOfLegs() int
}

type Dog struct{
	Name string
	Breed string
}

type Gorilla struct{
	Name string
	Breed string
	NumberOfTeeth int
}

func main(){
	dog := Dog{"Samson", "Shephard"}

	PrintInfo(&dog)
}

func PrintInfo(a Animal){
	fmt.Println("This animal says ,", a.Says(), "and has ", a.NumberOfLegs(), " legs");
}

func (d *Dog) Says() string{
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}