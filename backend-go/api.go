package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle this error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.Use(commonMiddleware)

	router.HandleFunc("/user", makeHTTPHandlerFunc(s.handleUser))
	router.HandleFunc("/user/{id}", makeHTTPHandlerFunc(s.handleGetUser))

	log.Println("server running on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func newAPIServer(
	listenAddr string,
) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) handleUser(
	w http.ResponseWriter, r *http.Request,
) error {
	if r.Method == "GET" {
		return s.handleGetUser(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetUser(
	w http.ResponseWriter, r *http.Request,
) error {
	return WriteJSON(w, http.StatusCreated, nil)
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
	return WriteJSON(w, http.StatusCreated, user)
}

func (s *APIServer) handleDeleteUser(
	w http.ResponseWriter, r *http.Request,
) error {
	return nil
}
