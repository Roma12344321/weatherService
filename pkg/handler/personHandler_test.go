package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"weatherService/pkg/service"
	"weatherService/pkg/service/mocks"
)

func TestHandler_createPerson(t *testing.T) {
	mockAuthService := new(mocks.AuthService)
	mockAuthService.On("Registration", mock.AnythingOfType("model.Person")).Return(1, nil)
	h := &Handler{
		service: &service.Service{
			AuthService: mockAuthService,
		},
	}
	r := gin.Default()
	r.POST("/registration", h.createPerson)
	personJSON := `{"username":"testuser","password":"password123"}`
	req, _ := http.NewRequest("POST", "/registration", bytes.NewBufferString(personJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"id":1}`, w.Body.String())
	mockAuthService.AssertExpectations(t)
}

func TestHandler_logIn(t *testing.T) {
	mockAuthService := new(mocks.AuthService)
	mockAuthService.On("GenerateToken", "testuser", "password123").Return("mock_token", nil)
	h := &Handler{
		service: &service.Service{
			AuthService: mockAuthService,
		},
	}
	r := gin.Default()
	r.POST("/login", h.logIn)
	personJSON := `{"username":"testuser","password":"password123"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(personJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"token":"mock_token"}`, w.Body.String())
	mockAuthService.AssertExpectations(t)
}
