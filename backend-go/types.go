package main

import (
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type APIServer struct {
	listenAddr string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    Gender    `json:"gender"`
	Age       int64     `json:"age"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRegistrationRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    Gender `json:"gender"`
	Age       int64  `json:"age"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func NewUser(firstName, email, lastName, password string, age int64, gender Gender) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Gender:    gender,
		Email:     email,
		Password:  string(encpw),
		CreatedAt: time.Now().Local().UTC(),
	}, nil
}
