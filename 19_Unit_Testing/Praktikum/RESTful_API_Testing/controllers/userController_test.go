package controllers

import (
	"Praktikum/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersControlleValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)

	err := GetUsersController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUserControlleValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := GetUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUserControlleInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := GetUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code) // Perbaiki status yang diharapkan menjadi 404.
}

func TestCreateUserControlleValid(t *testing.T) {
	user := models.User{
		Name:     "Dylee testing",
		Email:    "TestingEmail@email.com",
		Password: "Password",
	}
	requestData, _ := json.Marshal(user)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateUserControllerValid(t *testing.T) {
	user := models.User{
		Name:     "UpdateTest",
		Email:    "TestingUpdate@email.com",
		Password: "PasswordUpdateTest",
	}
	requestData, _ := json.Marshal(user)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/users/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("11")

	err := UpdateUserController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateUserControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := UpdateUserController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteUserControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("11")

	err := DeleteUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteUserControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/users/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := DeleteUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}


