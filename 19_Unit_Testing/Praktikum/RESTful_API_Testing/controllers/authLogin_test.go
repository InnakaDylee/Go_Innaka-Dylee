package controllers

import (
	"Praktikum/configs"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitTestDB() *echo.Echo {
	configs.ConnectDBTest()
	e := echo.New()
	return e
}

func TestLoginValid(t *testing.T) {
	loginUser := `{"Username": "testuser", "Password": "password123"}`
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(loginUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) 	
	rec := httptest.NewRecorder()
	c := e.NewContext(req,rec)
	
	err := Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestLoginInvalid(t *testing.T) {
	loginUser := `{"Username": "", "Password": ""}`
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(loginUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) 	
	rec := httptest.NewRecorder()
	c := e.NewContext(req,rec)
	
	err := Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}