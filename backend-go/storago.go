package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(*User) error
	DeleteUserById(int) error
	// UpdateUser(*User) error
	GetAllUsers() ([]*User, error)
	GetUserById(int) (*User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connectionStr := "user=postgres dbname=go password=password sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createUserTable()
}

func (s *PostgresStore) createUserTable() error {
	query := `create table if not exists users (
		id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		email varchar(100),
		gender varchar(6),
		age integer,
		password varchar(100),
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateUser(user *User) error {
	query := `insert into users 
	(first_name, last_name, email,gender,age,password, created_at)
	values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.db.Query(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Gender,
		user.Age,
		user.Password,
		user.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *PostgresStore) DeleteUserById(id int) error {
	_, err := s.db.Query("delete from users where id = $1", id)
	return err
}

func (s *PostgresStore) GetUserById(id int) (*User, error) {
	rows, err := s.db.Query("select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("User %d not found", id)
}

func (s *PostgresStore) GetAllUsers() ([]*User, error) {
	rows, err := s.db.Query("select * from users")
	if err != nil {
		return nil, err
	}

	Users := []*User{}
	for rows.Next() {
		User, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		Users = append(Users, User)
	}

	return Users, nil
}

func scanIntoUser(rows *sql.Rows) (*User, error) {
	User := new(User)
	err := rows.Scan(
		&User.ID,
		&User.FirstName,
		&User.LastName,
		&User.Email,
		&User.Age,
		&User.Password,
		&User.Gender,
		&User.CreatedAt,
	)

	return User, err
}
