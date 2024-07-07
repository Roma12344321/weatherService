package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
)

func TestPersonRepository_CreatePerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewPersonRepositoryImpl(sqlxDB)
	person := model.Person{
		Username: "testuser",
		Password: `test`,
	}
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO person(username, password) VALUES ($1,$2) RETURNING id`)).
		WithArgs(person.Username, person.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := repo.CreatePerson(person)
	assert.NoError(t, err, "CreatePerson should not return an error")
	assert.Equal(t, 1, id, "Expected person ID to be 1")
}

func TestPersonRepository_GetPersonByUsernameAndPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewPersonRepositoryImpl(sqlxDB)

	person := model.Person{
		Id:       1,
		Username: "testuser",
		Password: "password123",
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM person WHERE username=$1 AND password=$2`)).
		WithArgs(person.Username, person.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
			AddRow(person.Id, person.Username, person.Password))
	result, err := repo.GetPersonByUsernameAndPassword(person.Username, person.Password)
	assert.NoError(t, err, "GetPersonByUsernameAndPassword should not return an error")
	assert.Equal(t, person, result, "Expected person to match")
}
