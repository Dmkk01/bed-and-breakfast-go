package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "1234567890",
		Age:         25,
		BirthDate:   time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	log.Println(user)

	myMap := make(map[string]string)

	myMap["firstName"] = "John"

	log.Println(myMap)

	var mySlice []string

	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Johnn")
	mySlice = append(mySlice, "Adam")

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(number)

	log.Println(number[0:2])
}
