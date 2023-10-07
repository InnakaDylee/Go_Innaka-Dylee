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

func TestGetBooksControlleValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)

	err := GetBooksController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBookControlleValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := GetBookController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBookControlleInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := GetBookController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code) // Perbaiki status yang diharapkan menjadi 404.
}

func TestCreateBookControlleValid(t *testing.T) {
	book := models.Book{
		Judul:     "Testing Book",
		Penulis:    "Testing Penulis",
		Penerbit: "Testing Penerbit",
	}
	requestData, _ := json.Marshal(book)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateBookController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateBookControllerValid(t *testing.T) {
	book := models.Book{
		Judul:     "Testing Book Update",
		Penulis:    "Testing Penulis Update",
		Penerbit: "Testing Penerbit Update",
	}
	requestData, _ := json.Marshal(book)
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/books/", bytes.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("4")

	err := UpdateBookController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateBookControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodPut, "/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := UpdateBookController(c)
	fmt.Println(rec.Code)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteBookControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("4")

	err := DeleteBookController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteBookControllerInvalid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/books/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := DeleteBookController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}


