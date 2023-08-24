package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.Use(ResponseTypeMiddleware)

	router.HandleFunc("/user", makeHTTPHandlerFunc(s.handleUser))
	router.HandleFunc("/user/{id}", makeHTTPHandlerFunc(s.handleManageUser))

	log.Println("server running on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func newAPIServer(
	listenAddr string,
	storage Storage,
) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      storage,
	}
}

func (s *APIServer) handleUser(
	w http.ResponseWriter, r *http.Request,
) error {
	if r.Method == "GET" {
		return s.handleAllUsers(w, r)
	} else if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	} else {
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleManageUser(
	w http.ResponseWriter, r *http.Request,
) error {
	if r.Method == "GET" {
		return s.handleGetUserById(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	} else {
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleCreateUser(
	w http.ResponseWriter, r *http.Request,
) error {
	req := new(UserRegistrationRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	user, err := NewUser(req.FirstName, req.LastName, req.Email, req.Password, (req.Age), req.Gender)
	if err != nil {
		return err
	}
	if err := s.store.CreateUser(user); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusCreated, user)
}

func (s *APIServer) handleDeleteUser(
	w http.ResponseWriter, r *http.Request,
) error {
	id, err := getId(r)

	if err != nil {
		return err
	}

	if err := s.store.DeleteUserById(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusNoContent, map[string]string{"message": "user deleted successfully" + strconv.Itoa(id)})
}

func (s *APIServer) handleAllUsers(
	w http.ResponseWriter, r *http.Request,
) error {
	users, err := s.store.GetAllUsers()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleGetUserById(
	w http.ResponseWriter, r *http.Request,
) error {
	id, err := getId(r)

	if err != nil {
		return err
	}

	user, err := s.store.GetUserById(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)

}
