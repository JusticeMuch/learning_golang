package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Firstname string `json:"firstname"`
}

func main(){
	myJson := `[{"firstname" : "clark"},{"firstname": "jjjjj"}]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Panicln("error unmarshalling", err)
	}

	log.Printf("unmarshalled %v", unmarshalled)

	var mySLice []Person 

	mySLice = append(mySLice, Person{"KKKKK"})
	mySLice = append(mySLice, Person{"kdfkd"})

	newJson, err := json.MarshalIndent(mySLice, "", "		")
	if err != nil{
		log.Println("error marshalling ", err)
	}

	log.Println(string(newJson))

}