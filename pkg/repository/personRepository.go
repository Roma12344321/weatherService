package repository

import (
	"github.com/jmoiron/sqlx"
	"weatherService/pkg/model"
)

type PersonRepositoryImpl struct {
	db *sqlx.DB
}

func NewPersonRepositoryImpl(db *sqlx.DB) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{db: db}
}

func (r *PersonRepositoryImpl) CreatePerson(person model.Person) (int, error) {
	query := "INSERT INTO person(username, password) VALUES ($1,$2) RETURNING id"
	row := r.db.QueryRow(query, person.Username, person.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PersonRepositoryImpl) GetPersonByUsernameAndPassword(username, password string) (model.Person, error) {
	var person model.Person
	query := "SELECT * FROM person WHERE username=$1 AND password=$2"
	if err := r.db.Get(&person, query, username, password); err != nil {
		return person, err
	}
	return person, nil
}
