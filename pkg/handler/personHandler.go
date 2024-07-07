package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"weatherService/pkg/model"
)

const (
	authHeader = "Authorization"
	personCtx  = "person_id"
)

type inputUsernameAndPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) createPerson(c *gin.Context) {
	var inputData inputUsernameAndPassword
	if err := c.BindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	id, err := h.service.Registration(model.Person{Username: inputData.Username, Password: inputData.Password})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) logIn(c *gin.Context) {
	var person model.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	token, err := h.service.AuthService.GenerateToken(person.Username, person.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) personIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.service.AuthService.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(personCtx, userId)
}

func getPersonId(c *gin.Context) (int, error) {
	id, ok := c.Get(personCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
