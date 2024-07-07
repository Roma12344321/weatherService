package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
	"weatherService/pkg/repository/mocks"
	"weatherService/pkg/service"
)

func TestAuthService_Registration(t *testing.T) {
	mockPersonRepo := new(mocks.PersonRepository)
	authService := service.NewAuthServiceImpl(&repository.Repository{
		PersonRepository: mockPersonRepo,
	})
	person := model.Person{Username: "user", Password: "password"}
	hashedPassword := service.GeneratePasswordHash(person.Password)
	mockPersonRepo.On("CreatePerson", mock.MatchedBy(func(p model.Person) bool {
		return p.Username == person.Username && p.Password == hashedPassword
	})).Return(1, nil)
	id, err := authService.Registration(person)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockPersonRepo.AssertExpectations(t)
}

func TestAuthService_GenerateToken(t *testing.T) {
	mockPersonRepo := new(mocks.PersonRepository)
	authService := service.NewAuthServiceImpl(&repository.Repository{
		PersonRepository: mockPersonRepo,
	})
	username := "user"
	password := "password"
	hashedPassword := service.GeneratePasswordHash(password)
	person := model.Person{Id: 1, Username: username, Password: hashedPassword}
	mockPersonRepo.On("GetPersonByUsernameAndPassword", username, hashedPassword).Return(person, nil)
	token, err := authService.GenerateToken(username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockPersonRepo.AssertExpectations(t)
}

func TestAuthService_ParseToken(t *testing.T) {
	mockPersonRepo := new(mocks.PersonRepository)
	authService := service.NewAuthServiceImpl(&repository.Repository{
		PersonRepository: mockPersonRepo,
	})
	username := "user"
	password := "password"
	hashedPassword := service.GeneratePasswordHash(password)
	person := model.Person{Id: 1, Username: username, Password: hashedPassword}
	mockPersonRepo.On("GetPersonByUsernameAndPassword", username, hashedPassword).Return(person, nil)
	token, err := authService.GenerateToken(username, password)
	assert.NoError(t, err)
	id, err := authService.ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, person.Id, id)
	mockPersonRepo.AssertExpectations(t)
}
