package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "key"
)

type AuthServiceImpl struct {
	repo *repository.Repository
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"person_id"`
}

func NewAuthServiceImpl(repo *repository.Repository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) Registration(person model.Person) (int, error) {
	passwordHash := GeneratePasswordHash(person.Password)
	person.Password = passwordHash
	id, err := s.repo.PersonRepository.CreatePerson(person)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *AuthServiceImpl) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetPersonByUsernameAndPassword(username, GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthServiceImpl) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
