package main

import "math/rand"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int64  `json:"age"`
}

func NewUser(firstName, lastName string, age int) *User {
	return &User{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Age:       int64(age),
	}
}
