package controllers

import (
	"Praktikum/configs"
	"Praktikum/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	var books []models.Book
  
	if err := configs.DB.Find(&books).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all books",
	  "users":   books,
	})
}

func GetBookController(c echo.Context) error {
	var books []models.Book
	var book models.Book

	id, _ := strconv.Atoi(c.Param("id"))
	if err := configs.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	bookFound := configs.DB.First(&book,id)
	bookValue := configs.DB.Find(&books, id)
	if bookFound.Error == nil {
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"users":   bookValue.Value,
	  })
	}
  
	return c.JSON(http.StatusNotFound, map[string]interface{}{
	  "message": "book not found",
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	if err := configs.DB.Save(&book).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new book",
	  "book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	// your solution here
	var book models.Book
	var books []models.Book
	id, _ := strconv.Atoi(c.Param("id"))
  
	if err := configs.DB.Find(&books).Error ;err != nil{
	  return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"massage": "fail get book",
	  })
	}
	bookFound := configs.DB.First(&book,id)
	bookValue := configs.DB.Find(&books, id)
	if bookFound.Error == nil{
	  configs.DB.Delete(&books, id)
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "delete book success",
		"book": bookValue,
	  })
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
	  "massage": "book not found",
	})
}

func UpdateBookController(c echo.Context) error {
	// your solution here
	var book models.Book
	var edit models.Book
	c.Bind(&edit)
  
	id,_ := strconv.Atoi(c.Param("id"))
	
	if err := configs.DB.Find(&book).Error ;err != nil{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
		  "massage": "fail get book",
		})
	  }
	bookFound := configs.DB.First(&book,id)
	if bookFound.Error != nil{
		return c.JSON(http.StatusNotFound, map[string]interface{}{
		  "massage": "book not found",
		})
	}
	book.Judul = edit.Judul
	book.Penulis = edit.Penulis
	book.Penerbit = edit.Penerbit
  
	configs.DB.Save(&book)
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "massage": "update book success",
	  "book" : book,
	})
}