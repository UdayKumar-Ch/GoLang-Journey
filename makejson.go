package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Address string
}

func main() {

	var name string
	var address string
	fmt.Print("Enter Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Address: ")
	fmt.Scan(&address)

	var per Person
	per = Person{name, address}
	b, _ := json.Marshal(per)

	fmt.Println("JSON : ", string(b))

}
